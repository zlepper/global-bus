package global_bus

import (
	"context"
	"google.golang.org/protobuf/proto"
	"reflect"
)

type Context interface {
	context.Context
	Send(message proto.Message) error
}

type subscriptionCallbackFunc func(message interface{}, context Context) error

type subscriptionCallback struct {
	eventPath string
	callback  subscriptionCallbackFunc
}

type globalBusBuilder struct {
	// The transport configuration to use
	transport TransportConfiguration
	// The messages to subscribe to
	subscriptions map[string][]subscriptionCallback
	// Messages that should no longer be subscribed
	removedSubscriptions []string
	// The logger to use
	logger betterLogger
}

func getEventPath(m proto.Message) string {
	options := m.ProtoReflect().Descriptor().Options()
	if !proto.HasExtension(options, E_EventPath) {
		return ""
	}
	return proto.GetExtension(options, E_EventPath).(string)
}

func (b *globalBusBuilder) Subscribe(callback interface{}) GlobalBusCompleteBuilder {
	actualCallbackType := reflect.TypeOf(callback)
	if actualCallbackType.Kind() != reflect.Func {
		b.logger.Panic("Provided callback is not a function")
	}

	if actualCallbackType.NumIn() > 2 {
		b.logger.Panic("Provided callback takes more than 2 arguments")
	}

	if actualCallbackType.NumIn() == 0 {
		b.logger.Panic("Provided callback doesn't take any arguments")
	}

	if actualCallbackType.NumOut() > 1 {
		b.logger.Panic("Provided callback has to many return values")
	}

	if actualCallbackType.NumOut() == 1 {
		out := actualCallbackType.Out(0)
		if out != reflect.TypeOf((*error)(nil)).Elem() {
			b.logger.Panic("Provided callback has a return argument, but it is not of the 'error' type")
		}
	}

	if actualCallbackType.NumIn() == 2 {
		contextArg := actualCallbackType.In(1)
		if contextArg != reflect.TypeOf((*Context)(nil)).Elem() {
			b.logger.Panic("Second argument to provided callback is not of the type `global_bus.Context`")
		}
	}

	messageArg := actualCallbackType.In(0)

	if !reflect.PtrTo(messageArg).Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		b.logger.Panic("First argument of callback is not a protobuf message")
	}

	instance := reflect.New(messageArg).Interface().(proto.Message)

	eventPath := getEventPath(instance)
	if eventPath == "" {
		b.logger.Panic("First argument of callback doesn't provide an event_path option")
	}

	var callbackWrapper subscriptionCallbackFunc
	var actualCallbackValue = reflect.ValueOf(callback)

	if actualCallbackType.NumOut() == 1 {
		if actualCallbackType.NumIn() == 2 {
			callbackWrapper = contextArgumentWithErrorCallbackWrapper(actualCallbackValue)
		} else {
			callbackWrapper = singleArgumentWithErrorCallbackWrapper(actualCallbackValue)
		}
	} else {
		if actualCallbackType.NumIn() == 2 {
			callbackWrapper = contextArgumentNoErrorCallbackWrapper(actualCallbackValue)
		} else {
			callbackWrapper = singleArgumentNoErrorCallbackWrapper(actualCallbackValue)
		}
	}

	cb := subscriptionCallback{
		eventPath: eventPath,
		callback:  callbackWrapper,
	}

	existing, exists := b.subscriptions[eventPath]
	if exists {
		existing = append(existing, cb)
	} else {
		existing = []subscriptionCallback{cb}
	}
	b.subscriptions[eventPath] = existing

	return b
}

func singleArgumentNoErrorCallbackWrapper(callback reflect.Value) subscriptionCallbackFunc {
	return func(message interface{}, context Context) error {
		callback.Call([]reflect.Value{reflect.ValueOf(message)})
		return nil
	}
}

func singleArgumentWithErrorCallbackWrapper(callback reflect.Value) subscriptionCallbackFunc {
	return func(message interface{}, context Context) error {
		result := callback.Call([]reflect.Value{reflect.ValueOf(message)})
		return result[0].Interface().(error)
	}
}

func contextArgumentNoErrorCallbackWrapper(callback reflect.Value) subscriptionCallbackFunc {
	return func(message interface{}, context Context) error {
		callback.Call([]reflect.Value{reflect.ValueOf(message), reflect.ValueOf(context)})
		return nil
	}
}

func contextArgumentWithErrorCallbackWrapper(callback reflect.Value) subscriptionCallbackFunc {
	return func(message interface{}, context Context) error {
		result := callback.Call([]reflect.Value{reflect.ValueOf(message), reflect.ValueOf(context)})
		return result[0].Interface().(error)
	}
}

func (b *globalBusBuilder) Unsubscribe(sample proto.Message) GlobalBusCompleteBuilder {
	eventPath := getEventPath(sample)

	if eventPath == "" {
		b.logger.Panic("Provided sample is not a proper event message. It's missing the event_path option.")
	}

	b.removedSubscriptions = append(b.removedSubscriptions, eventPath)

	return b
}

func (b *globalBusBuilder) Start() (GlobalBusInstance, error) {
	panic("implement me")
}

func (b *globalBusBuilder) WithTransport(transport TransportConfiguration) GlobalBusCompleteBuilder {
	b.transport = transport
	return b
}

type GlobalBusBuilderTransportNext interface {
	// Configured GlobalBus to use the specified transport
	WithTransport(transport TransportConfiguration) GlobalBusCompleteBuilder
}

type GlobalBusCompleteBuilder interface {

	// Subscribes to the specified event type and calls the
	// callback automatically. The callback has to be a function
	// of either of the following signatures:
	//      func(event MyEvent) error
	//      func(event MyEvent)
	//      func(event MyEvent, context global_bus.Context) error
	//      func(event MyEvent, context global_bus.Context)
	//
	// This method panics if the provided callback is not a function
	// with any of the above signatures
	//
	// In other news: I'm really excited for proper generics in this language :D
	Subscribe(callback interface{}) GlobalBusCompleteBuilder
	// Unsubscribes to the specified kind of messages. A sample
	// instance has to be provided so Global Bus can know what to ignore
	// Any messages that are still queued of this kind we be thrown away when
	// they arrive.
	Unsubscribe(sample proto.Message) GlobalBusCompleteBuilder

	// Actually starts Global Bus. An error is returned if the
	// bus itself failed to start
	Start() (GlobalBusInstance, error)
}

type GlobalBusInstance interface{}

func CreateBuilder() GlobalBusBuilderTransportNext {
	return &globalBusBuilder{
		transport:            nil,
		subscriptions:        make(map[string][]subscriptionCallback),
		removedSubscriptions: make([]string, 0),
		logger:               betterLogger{StdLogger{}},
	}
}
