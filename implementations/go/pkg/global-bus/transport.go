package global_bus

type TransportCallback func(packet TransportPacket)

type Transport interface {
	Stop() error
	// Should add a new subscription for the specified eventPath
	AddSubscription(eventPath string)
	// Sets the handler that should be called whenever a new message is available
	OnNewPacket(packetCallback TransportCallback)
	// Should remove any subscriptions for the specified eventPath.
	// Any existing messages coming in of this type are dealt with in the framework
	// the transport doesn't need to handle that.
	RemoveSubscription(eventPath string)
}

type TransportConfiguration interface {
	// Should create a new instance of the transport for communication
	CreateInstance() (Transport, error)
}

type TransportPacket interface {
	// Should get the actual underlying message bytes
	GetData() []byte
	// Should remove the message from it's current queue. Often called an "Ack" operation
	MarkAsDone() error
	// Should put the message back in queue somehow if possible.
	PutBackInQueue() error
	// Should move the message into the error queue for manual inspection later on
	MoveToErrorQueue() error
}
