package application

import (
	"context"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/aggregate"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/event"
)

type UpdateProduct struct {
	ID          string  `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

func (svc ProductApplicationServiceImpl) UpdateProduct(ctx context.Context, input core.Event) core.Event {
	cmd := &UpdateProduct{}
	input.Data(cmd)

	entity := cevixe.Entity(ctx, "Product", cmd.ID)
	if entity == nil {
		return cevixe.NewBusinessEvent(ctx, &core.DomainEntityNotFound{Type: "Product", ID: cmd.ID})
	}

	entityState := &aggregate.Product{}
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

	return cevixe.NewDomainEvent(ctx, entity, newEvent, newState)
}
