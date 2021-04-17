package application

import (
	"context"
	"functions/product/pkg/domain/aggregate"
	"functions/product/pkg/domain/command"
	domainEvent "functions/product/pkg/domain/event"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
)

func (svc ProductApplicationServiceImpl) CreateProduct(ctx context.Context, event core.Event) core.Event {
	cmd := &command.CreateProduct{}
	event.Payload(cmd)

	newEvent := domainEvent.ProductCreated{
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
	}

	newState := aggregate.Product{
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
	}

	return cevixe.NewEvent(ctx, nil, newEvent, newState)
}
