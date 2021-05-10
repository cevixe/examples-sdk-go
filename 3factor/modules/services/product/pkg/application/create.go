package application

import (
	"context"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/aggregate"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/event"
)

type CreateProduct struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

func (svc ProductApplicationServiceImpl) CreateProduct(ctx context.Context, input core.Event) core.Event {
	cmd := &CreateProduct{}
	input.Data(cmd)

	newEvent := &event.ProductCreated{
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
	}

	newState := &aggregate.Product{
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
	}

	return cevixe.NewDomainEvent(ctx, nil, newEvent, newState)
}
