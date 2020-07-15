package global_bus

import (
	"context"
	"errors"
	"google.golang.org/protobuf/proto"
	"reflect"
	"sync"
)

type Context interface {
	context.Context
	Send(message proto.Message) error
}

type messageContext struct {
	context.Context
	queuedMessages []proto.Message
	lock           *sync.Mutex
}

var (
	ErrInvalidEventMessage = errors.New("invalid event message")
)

func (m *messageContext) Send(message proto.Message) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	eventPath := getEventPath(message)
	if eventPath == "" {
		return ErrInvalidEventMessage
	}

	m.queuedMessages = append(m.queuedMessages, message)
	return nil
}

type subscriptionCallbackFunc func(message interface{}, context Context) error

type subscriptionCallback struct {
	eventPath string
	callback  subscriptionCallbackFunc
	eventType reflect.Type
}
