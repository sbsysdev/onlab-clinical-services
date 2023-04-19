package shareddomain

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type DomainEvent string

type Event interface {
	EventName() DomainEvent
	EventOccurredAt() time.Time
}

type PublishDomainEvent func(Event) error

type SubscribeDomainEvent func(context.Context, DomainEvent) (<-chan *message.Message, error)
