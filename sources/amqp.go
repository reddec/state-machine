package sources

import (
	"context"
	"github.com/reddec/fluent-amqp"
	"github.com/reddec/state-machine/machine"
	"github.com/streadway/amqp"
)

const AMQPMessageKey = "message"

func MakeAMQPHandler(vm *machine.StateMachine) fluent.TransactionHandlerFunc {
	return func(ctx context.Context, msg amqp.Delivery) error {
		id := msg.CorrelationId
		if id == "" {
			id = msg.MessageId
		}
		return vm.Emit(context.WithValue(ctx, AMQPMessageKey, msg), id, msg.Body)
	}
}

func MakeAMQPHandlerWithRetry(vm *machine.StateMachine, retryQueue fluent.Requeue) fluent.TransactionHandlerFunc {
	return func(ctx context.Context, msg amqp.Delivery) error {
		id := msg.CorrelationId
		if id == "" {
			id = msg.MessageId
		}
		err := vm.Emit(context.WithValue(ctx, AMQPMessageKey, msg), id, msg.Body)
		if err != nil {
			return retryQueue.Requeue(&msg)
		}
		return nil
	}
}
