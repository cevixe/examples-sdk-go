package application

import (
	"context"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/aggregate"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/event"
	"reflect"
)

type DeleteProduct struct {
	ID string `json:"id,omitempty"`
}

func (svc ProductApplicationServiceImpl) DeleteProduct(ctx context.Context, input core.Event) core.Event {
	cmd := &DeleteProduct{}
	input.Payload(cmd)

	entityState := &aggregate.Product{}
	entity := cevixe.Entity(ctx, reflect.TypeOf(entityState).Name(), cmd.ID)

	return cevixe.NewEvent(ctx, entity, &event.ProductDeleted{}, nil)
}
