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

	cevixeEvents := make(map[string]core.Event)
	for _, item := range input.Records {
		cevixeEvent := dynamo.MapDynamoEventRecordToCevixeEvent(ctx, item)
		source := cevixeEvent.Source().Type() + "/" + cevixeEvent.Source().ID()
		cevixeEvents[source] = cevixeEvent
	}

	cevixeEntities := make([]core.Entity, 0, len(cevixeEvents))
	for _, value := range cevixeEvents {
		payload := &map[string]interface{}{}
		source := value.Source()
		source.State(payload)
		cevixeEntities = append(cevixeEntities, source)
	}

	stateStore := ctx.Value(cevixe.CevixeStateStore).(core.StateStore)
	stateStore.UpdateState(ctx, cevixeEntities)
	return nil
}

func main() {
	ctx := runtime.NewContext()
	lambda.StartWithContext(ctx, handler)
}
