package global_bus

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
)

const (
	te = "testEvent"
)

type TestContext struct {
	context.Context
	SendMessages []proto.Message
}

func (c *TestContext) Send(message proto.Message) error {
	c.SendMessages = append(c.SendMessages, message)
	return nil
}

var ev = MyValidTestEvent{}
var ctx Context = &TestContext{context.Background(), []proto.Message{}}

func TestSubscribeCreatesCorrectFunction_ContextArgWithError(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	called := false

	builder.WithTransport(nil).
		Subscribe(func(event MyValidTestEvent, context Context) error {
			called = true
			return errors.New("test")
		})

	cb := builder.subscriptions[te][0]

	assert.Equal(t, te, cb.eventPath)

	err := cb.callback(ev, ctx)

	assert.Error(t, err)
	assert.True(t, called)
}

func TestProducesProperErrors(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError string
		value         interface{}
	}{
		{
			"Not a func",
			"Provided callback is not a function",
			42,
		},
		{
			"Too many args",
			"Provided callback doesn't take exactly 2 arguments",
			func(ev MyValidTestEvent, context Context, something int) error {
				return nil
			},
		},
		{
			"Too few args 0",
			"Provided callback doesn't take exactly 2 arguments",
			func() error {
				return nil
			},
		},
		{
			"Too few args 1",
			"Provided callback doesn't take exactly 2 arguments",
			func(ev MyValidTestEvent) error {
				return nil
			},
		},
		{
			"Too many out",
			"doesn't have exactly 1 return type",
			func(ev MyValidTestEvent, context Context) (int, error) {
				return 0, nil
			},
		},
		{
			"Out is not an error type",
			"Provided callback has a return argument, but it is not of the 'error' type",
			func(ev MyValidTestEvent, context Context) int {
				return 0
			},
		},
		{
			"Context argument is not a context",
			"Second argument to provided callback is not of the type `global_bus.Context`",
			func(ev MyValidTestEvent, s int) error {
				return nil
			},
		},
		{
			"Event arg is not a proto event",
			"First argument of callback is not a protobuf message",
			func(s int, context Context) error {
				return nil
			},
		},
		{
			"Provided protobuf message doesn't have event_path option",
			"First argument of callback doesn't provide an event_path option",
			func(ev MyInvalidTestEvent, context Context) error {
				return nil
			},
		},
	}

	builder := CreateBuilder().(*globalBusBuilder)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.PanicsWithValue(t, testCase.expectedError, func() {
				builder.Subscribe(testCase.value)
			})
		})
	}
}

func TestGlobalBusBuilder_Subscribe_ToSameEventTwice(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	builder.Subscribe(func(ev MyValidTestEvent, context Context) error {
		return nil
	})
	builder.Subscribe(func(ev MyValidTestEvent, context Context) error {
		return nil
	})

	subs := builder.subscriptions[te]

	assert.Len(t, subs, 2)
	assert.NotEqual(t, subs[0], subs[1])
}

func TestGlobalBusBuilder_Unsubscribe_CanRemoveSubscription(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	builder.Unsubscribe(&MyValidTestEvent{})

	assert.Equal(t, []string{te}, builder.removedSubscriptions)
}

func TestGlobalBusBuilder_Unsubscribe_PanicsOnInvalid(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	assert.PanicsWithValue(t, "Provided sample is not a proper event message. It's missing the event_path option.", func() {
		builder.Unsubscribe(&MyInvalidTestEvent{})
	})

	assert.Empty(t, builder.removedSubscriptions)
}
