package global_bus

import (
	"context"
	"google.golang.org/protobuf/proto"
	"reflect"
	"sync"
)

type GlobalBusInstance struct {
	logger        betterLogger
	workerManager *workerManager
	context       context.Context
}

func (gb *GlobalBusInstance) startGlobalBusInstance(configuration *globalBusBuilder) error {
	gb.logger = configuration.logger
	// Check if shutdown has already been requested, in which case there
	// is no point in starting the server
	if configuration.context.Err() != nil {
		gb.logger.Debug("Context has been cancelled already, skipping Global Bus initialization")
		return nil
	}

	gb.logger.Debug("Creating transport instance")

	// Start the actual transport
	transport, err := configuration.transport.CreateInstance()
	if err != nil {
		gb.logger.Error("Failed to start transport: %v", err)
		return err
	}

	gb.logger.Debug("Created transport instance")

	transport.OnNewPacket(gb.handleTransportPacket(configuration.subscriptions, configuration.removedSubscriptions))

	for eventPath := range configuration.subscriptions {
		gb.logger.Debug("Adding subscription '%s'", eventPath)
		transport.AddSubscription(eventPath)
	}

	for _, eventPath := range configuration.removedSubscriptions {
		gb.logger.Debug("Removing subscription '%s'", eventPath)
		transport.RemoveSubscription(eventPath)
	}

	gb.logger.Info("Started GlobalBus instance")

	go func() {
		gb.logger.Debug("Global bus is waiting for a context cancellation to shutdown.")
		<-configuration.context.Done()
		gb.logger.Debug("Global Bus context was cancelled, waiting for workers to exit")
		gb.workerManager.workersFinished.Wait()
		gb.logger.Debug("All Global Bus worker has stopped. Stopping transport")
		err := transport.Stop()
		if err != nil {
			gb.logger.Error("Failed to properly stop transport: %v", err)
		} else {
			gb.logger.Info("Global Bus has been stopped")
		}
	}()

	return nil
}

func stringArrayContains(haystack []string, needle string) bool {
	for _, wheat := range haystack {
		if wheat == needle {
			return true
		}
	}

	return false
}

func (gb *GlobalBusInstance) handleTransportPacket(subscriptions map[string][]subscriptionCallback, removedSubscriptions []string) TransportCallback {
	return func(packet TransportPacket) {
		var message EventPackage
		err := proto.Unmarshal(packet.GetData(), &message)

		if err != nil {
			gb.logger.Error("Failed to unmarshal event packet: %v", err)
			gb.movePacketToErrorQueue(packet)
			return
		}

		if stringArrayContains(removedSubscriptions, message.EventPath) {
			gb.logger.Debug("Got message '%s' that has been removed from subscriptions. This was probably a previously stored message. Skipping...", message.EventPath)
			return
		}

		callbacks, exists := subscriptions[message.EventPath]
		if !exists {
			gb.logger.Error("Got message with no known handlers '%s'. Did you forget to unsubscribe?", message.EventPath)
			gb.movePacketToErrorQueue(packet)
			return
		} else {
			gb.logger.Debug("Got message '%s', with %d callbacks", message.EventPath, len(callbacks))

			var hadError bool

			var wg sync.WaitGroup
			wg.Add(len(callbacks))
			for _, callback := range callbacks {
				go func(cb subscriptionCallback) {

					gb.workerManager.process(func() {

						value := reflect.New(cb.eventType).Interface().(proto.Message)
						proto.Unmarshal(message.EventData, value)

						ctx := messageContext{
							Context:        gb.context,
							queuedMessages: make([]proto.Message, 0),
							lock:           &sync.Mutex{},
						}

						err := cb.callback(value, &ctx)
						if err != nil {
							hadError = true
							gb.logger.Error("Event handler failed: %v. ", err)
						}

					})

					wg.Done()
				}(callback)
			}
			wg.Wait()

		}

	}
}

func (gb *GlobalBusInstance) movePacketToErrorQueue(packet TransportPacket) {
	err := packet.MoveToErrorQueue()
	if err != nil {
		gb.logger.Critical("Failed to move failed message to error queue: %v", err)
		// If we cannot move the message into the error queue, then we'll refuse to mark it as done
		return
	}
	err = packet.MarkAsDone()
	if err != nil {
		gb.logger.Error("Failed to remove errored message from normal queue: %v", err)
	}
}
