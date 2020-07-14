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

func TestSubscribeCreatesCorrectFunction_SingleArgNoError(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	called := false

	builder.WithTransport(nil).
		Subscribe(func(event MyValidTestEvent) {
			called = true
		})

	cb := builder.subscriptions[te][0]

	assert.Equal(t, te, cb.eventPath)

	err := cb.callback(ev, ctx)

	assert.NoError(t, err)
	assert.True(t, called)
}

func TestSubscribeCreatesCorrectFunction_SingleArgWithError(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	called := false

	builder.WithTransport(nil).
		Subscribe(func(event MyValidTestEvent) error {
			called = true
			return errors.New("test")
		})

	cb := builder.subscriptions[te][0]

	assert.Equal(t, te, cb.eventPath)

	err := cb.callback(ev, ctx)

	assert.Error(t, err)
	assert.True(t, called)
}

func TestSubscribeCreatesCorrectFunction_ContextArgNoError(t *testing.T) {
	builder := CreateBuilder().(*globalBusBuilder)

	called := false

	builder.WithTransport(nil).
		Subscribe(func(event MyValidTestEvent, context Context) {
			called = true
		})

	cb := builder.subscriptions[te][0]

	assert.Equal(t, te, cb.eventPath)

	err := cb.callback(ev, ctx)

	assert.NoError(t, err)
	assert.True(t, called)
}

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
			"Provided callback takes more than 2 arguments",
			func(ev MyValidTestEvent, context Context, something int) {},
		},
		{
			"Too few args",
			"Provided callback doesn't take any arguments",
			func() {},
		},
		{
			"Too many out",
			"Provided callback has to many return values",
			func(ev MyValidTestEvent) (int, error) {
				return 0, nil
			},
		},
		{
			"Out is not an error type",
			"Provided callback has a return argument, but it is not of the 'error' type",
			func(ev MyValidTestEvent) int {
				return 0
			},
		},
		{
			"Context argument is not a context",
			"Second argument to provided callback is not of the type `global_bus.Context`",
			func(ev MyValidTestEvent, s int) {},
		},
		{
			"Event arg is not a proto event",
			"First argument of callback is not a protobuf message",
			func(s int) {},
		},
		{
			"Provided protobuf message doesn't have event_path option",
			"First argument of callback doesn't provide an event_path option",
			func(ev MyInvalidTestEvent) {},
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

	builder.Subscribe(func(ev MyValidTestEvent) {})
	builder.Subscribe(func(ev MyValidTestEvent) {})

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
