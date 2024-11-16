package consumer

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// Handler function for consumer group
type Handler func(ctx context.Context, msg *sarama.ConsumerMessage) error

// GroupHandler represents consumer group
type GroupHandler struct {
	msgHandler Handler
}

// NewGroupHandler return new Consumer Group
func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

// Setup is run at the start of a new session before ConsumeClaim is called
func (c *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup runs at the end of the session life after all ConsumeClaim goroutines have completed
func (c *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim should start the ConsumerGroupClaim() message consumer loop.
// After the Messages() channel is closed, the handler should finish processing
func (c *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel wss closed\n")
				return nil
			}

			log.Printf("message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)

			err := c.msgHandler(session.Context(), message)
			if err != nil {
				log.Printf("error handling message: %v", err)
				continue
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			log.Printf("session conext done")
			return nil
		}
	}
}
