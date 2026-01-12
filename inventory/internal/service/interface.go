package service

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/model"
	inventoryv1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/inventory/v1"
)

type Service interface {
	GetPart(ctx context.Context, uuid string) (*model.Part, error)
	ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error)

	inventoryv1.UnimplementedInventoryServiceServer
}
