package repository

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/model"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, uuid string) (*model.Part, error)
	ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error)
	CreatePart(ctx context.Context, part *model.Part) error
	UpdatePart(ctx context.Context, part *model.Part) error
	DeletePart(ctx context.Context, uuid string) error
}
