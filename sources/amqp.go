package sources

import (
	"context"
	"github.com/reddec/fluent-amqp"
	"github.com/reddec/state-machine/machine"
	"github.com/streadway/amqp"
)

func MakeAMQPHandler(vm *machine.StateMachine) fluent.TransactionHandlerFunc {
	return func(ctx context.Context, msg amqp.Delivery) error {
		id := msg.CorrelationId
		if id == "" {
			id = msg.MessageId
		}
		return vm.Emit(context.WithValue(ctx, "message", msg), id, msg.Body)
	}
}
