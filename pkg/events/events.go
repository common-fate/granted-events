package events

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	eTypes "github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/common-fate/granted-events/pkg/types"
)

//EventSender provides methods to submit events to a granted event bridge
type EventSender struct {
	client      *eventbridge.Client
	eventBusArn string
	source      string
}
type Opts struct {
	EventBusARN string
	//SourceName is a descriptor for the service sending the events e.g "GRANTED_ACCESS_HANDLER"
	SourceName string
}

//New creates a new EventSender
func New(ctx context.Context, opts Opts) (*EventSender, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &EventSender{
		client:      eventbridge.NewFromConfig(cfg),
		eventBusArn: opts.EventBusARN,
		source:      opts.SourceName,
	}, nil
}

type EventTyper interface {
	EventType() types.EventType
}

func (e *EventSender) Put(ctx context.Context, detail EventTyper) error {
	d, err := json.Marshal(detail)
	if err != nil {
		return err
	}

	entry := eTypes.PutEventsRequestEntry{
		EventBusName: &e.eventBusArn,
		Detail:       aws.String(string(d)),
		DetailType:   aws.String(string(detail.EventType())),
		Source:       &e.source,
	}
	res, err := e.client.PutEvents(ctx, &eventbridge.PutEventsInput{
		Entries: []eTypes.PutEventsRequestEntry{entry},
	})
	if err != nil {
		return err
	}
	if res.FailedEntryCount != 0 {
		return fmt.Errorf("failed to send error with code: %s, error: %s", *res.Entries[0].ErrorCode, *res.Entries[0].ErrorMessage)
	}
	return nil
}
