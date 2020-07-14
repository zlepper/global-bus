package global_bus

type Transport interface {
}

type TransportConfiguration interface {
	// Should create a new instance of the transport for communication
	CreateInstance() (Transport, error)
}
