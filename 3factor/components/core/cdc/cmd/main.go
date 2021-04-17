package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cevixe/aws-sdk-go/aws/integration/dynamo"
	"github.com/cevixe/aws-sdk-go/runtime"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
)

func handler(ctx context.Context, input events.DynamoDBEvent) error {
	item := input.Records[0]
	event := dynamo.MapDynamoEventRecordToCevixeEvent(ctx, item)
	eventBus := ctx.Value(cevixe.CevixeEventBus).(core.EventBus)
	eventBus.PublishEvent(ctx, event)
	return nil
}

func main() {
	ctx := runtime.NewContext()
	lambda.StartWithContext(ctx, handler)
}
