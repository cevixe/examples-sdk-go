package application

import (
	"context"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/aggregate"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/event"
	"reflect"
)

type UpdateProduct struct {
	ID          string  `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

func (svc ProductApplicationServiceImpl) UpdateProduct(ctx context.Context, input core.Event) core.Event {
	cmd := &UpdateProduct{}
	input.Payload(cmd)

	entityState := &aggregate.Product{}
	entity := cevixe.Entity(ctx, reflect.TypeOf(entityState).Name(), cmd.ID)
	entity.State(entityState)

	newDescription := entityState.Description
	newPrice := entityState.Price

	if cmd.Description != "" {
		newDescription = cmd.Description
	}

	if cmd.Price > 0 {
		newPrice = cmd.Price
	}

	newState := &aggregate.Product{
		Name:        entityState.Name,
		Description: newDescription,
		Price:       newPrice,
	}

	newEvent := &event.ProductUpdated{
		Description: cmd.Description,
		Price:       cmd.Price,
	}

	return cevixe.NewEvent(ctx, entity, newEvent, newState)
}
