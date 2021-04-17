package application

import (
	"context"
	"github.com/cevixe/core-sdk-go/core"
)

type ProductApplicationService interface {
	CreateProduct(ctx context.Context, event core.Event) core.Event
	UpdateProduct(ctx context.Context, event core.Event) core.Event
}

type ProductApplicationServiceImpl struct {
	ProductApplicationService
}

func New() ProductApplicationService {
	return &ProductApplicationServiceImpl{}
}
