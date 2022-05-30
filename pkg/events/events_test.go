package events

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/common-fate/granted-events/pkg/types"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

func skipIfAwsCredsMissing(ctx context.Context, t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Skip(err)
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		t.Skip(err)
	}
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		t.Skip(err)
	}
	if !creds.HasKeys() || creds.Expired() {
		t.Skip("aws creds expired or missing")
	}
}

func TestPutEvent(t *testing.T) {
	ctx := context.Background()
	skipIfAwsCredsMissing(ctx, t)

	type config struct {
		EventBusARN string `env:"EVENT_BUS_ARN,required"`
	}
	var c config
	err := envconfig.Process(ctx, &c)
	if err != nil {
		t.Error(err)
	}

	e, err := New(ctx, Opts{EventBusARN: c.EventBusARN, SourceName: "granted-events-test"})
	if err != nil {
		t.Error(err)
	}

	ge := types.GrantEvent{
		Event: types.Event{
			Type: types.GrantStart,
		},
		GrantID: "abcd",
	}
	err = e.Put(ctx, ge)
	if err != nil {
		t.Error(err)
	}
}
