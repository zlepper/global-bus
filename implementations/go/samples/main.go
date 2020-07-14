package main

import (
	global_bus "github.com/zlepper/global-bus/implementations/go/pkg/global-bus"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

func main() {
	// proto.GetExtension(m.ProtoReflect().Descriptor().Options(), global_bus.E_EventPath)
	m := MyEvent{}

	before := time.Now()
	for i := 0; i < 10_000_000; i++ {
		_ = proto.GetExtension(m.ProtoReflect().Descriptor().Options(), global_bus.E_EventPath)
	}

	after := time.Now()

	diff := after.Sub(before)

	log.Printf("Time taken: %s", diff.String())

	//if e == nil {
	//	log.Printf("Nothing was found :(")
	//} else {
	//	s := e.(string)
	//	log.Printf("Found value: %s", s)
	//}
}
