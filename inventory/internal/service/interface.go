package service

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/model"
)

type Service interface {
	GetPart(ctx context.Context, uuid string) (*model.Part, error)
	ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error)
}
