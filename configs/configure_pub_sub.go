package configs

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/uuid"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// Configure Pub Sub
var pubSub = gochannel.NewGoChannel(
	gochannel.Config{},
	watermill.NewStdLogger(false, false),
)

// Publish Domain Event
func PublishDomainEvent(event shareddomain.Event) error {
	bytes, jsonErr := json.Marshal(event)

	if jsonErr != nil {
		return jsonErr
	}

	publishErr := pubSub.Publish(string(event.EventName()), message.NewMessage(uuid.NewString(), bytes))

	if publishErr != nil {
		return publishErr
	}

	return nil
}

// Subscribe Domain Event
func SubscribeDomainEvent(ctx context.Context, domainEvent shareddomain.DomainEvent) (<-chan *message.Message, error) {
	return pubSub.Subscribe(ctx, string(domainEvent))
}
