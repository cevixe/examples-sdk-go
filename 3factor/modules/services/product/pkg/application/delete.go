package application

import (
	"context"
	"github.com/cevixe/core-sdk-go/cevixe"
	"github.com/cevixe/core-sdk-go/core"
	"github.com/cevixe/examples-sdk-go/3factor/services/product/pkg/domain/event"
)

type DeleteProduct struct {
	ID string `json:"id,omitempty"`
}

func (svc ProductApplicationServiceImpl) DeleteProduct(ctx context.Context, input core.Event) core.Event {
	cmd := &DeleteProduct{}
	input.Data(cmd)

	entity := cevixe.Entity(ctx, "Product", cmd.ID)

	if entity == nil {
		return cevixe.NewBusinessEvent(ctx, &core.DomainEntityNotFound{Type: "Product", ID: cmd.ID})
	} else {
		return cevixe.NewDomainEvent(ctx, entity, &event.ProductDeleted{}, nil)
	}
}
