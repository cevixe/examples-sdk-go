package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cevixe/aws-sdk-go/aws/integration/dynamo"
	"github.com/cevixe/aws-sdk-go/aws/model"
	"github.com/cevixe/aws-sdk-go/runtime"
	"github.com/cevixe/aws-sdk-go/util"
	"github.com/cevixe/core-sdk-go/core"
	"sync"
)

const gqlQuery string = `
mutation(
	$Transaction: ID!
	$SourceId: ID!
	$SourceType: String!
	$SourceOwner: String!
	$EventAuthor: String!
	$EventType: String!
	$EventObject: AWSJSON!
) {
	publishEvent(
		transaction: $Transaction
		sourceId: $SourceId
		sourceType: $SourceType
		sourceOwner: $SourceOwner
		eventAuthor: $EventAuthor
		eventType: $EventType
		eventObject: $EventObject
	) {
		transaction
		eventId
		eventType
		eventTime
		eventAuthor
		eventPayload
		sourceId
		sourceType
		sourceTime
		sourceOwner
		sourceState
	}
}
`

func generateRequest(event core.Event) *model.GraphqlRequest {

	payload := &map[string]interface{}{}
	event.Payload(payload)

	state := &map[string]interface{}{}
	event.Source().State(state)

	return &model.GraphqlRequest{
		Query: gqlQuery,
		Variables: map[string]interface{}{
			"Transaction": event.Transaction(),
			"SourceId":    event.Source().ID(),
			"SourceType":  event.Source().Type(),
			"SourceOwner": event.Source().Owner(),
			"EventAuthor": event.Author(),
			"EventType":   event.Type(),
			"EventObject": util.MarshalJsonString(map[string]interface{}{
				"transaction":  event.Transaction(),
				"eventId":      event.ID(),
				"eventType":    event.Type(),
				"eventTime":    event.Time(),
				"eventAuthor":  event.Author(),
				"eventPayload": payload,
				"sourceId":     event.Source().ID(),
				"sourceType":   event.Source().Type(),
				"sourceTime":   event.Source().Time(),
				"sourceOwner":  event.Source().Owner(),
				"sourceState":  state,
			}),
		},
	}
}

func handler(ctx context.Context, input events.DynamoDBEvent) error {

	cevixeEvents := make([]core.Event, 0, len(input.Records))
	for _, item := range input.Records {
		cevixeEvent := dynamo.MapDynamoEventRecordToCevixeEvent(ctx, item)
		cevixeEvents = append(cevixeEvents, cevixeEvent)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(cevixeEvents))
	for _, item := range cevixeEvents {
		go asyncGraphqlCall(ctx, item, wg)
	}
	wg.Wait()
	return nil
}

func asyncGraphqlCall(ctx context.Context, event core.Event, wg *sync.WaitGroup) {
	request := generateRequest(event)
	awsContext := ctx.Value(model.AwsContext).(*model.Context)
	response, err := awsContext.AwsGraphGateway.ExecuteGraphql(ctx, request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: \n%s", util.MarshalJsonString(response))
	if len(response.Errors) > 0 {
		panic(fmt.Errorf("invalid gql request:\n%v", response.Errors))
	}
	wg.Done()
}

func main() {
	ctx := runtime.NewContext()
	lambda.StartWithContext(ctx, handler)
}
