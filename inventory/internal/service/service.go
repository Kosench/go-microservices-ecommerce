package service

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/model"
	inventoryv1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/inventory/v1"
)

type Service interface {
	GetPart(ctx context.Context, uuid string) (*inventoryv1.Part, error)
	ListParts(ctx context.Context, filter *inventoryv1.PartsFilter) ([]*model.Part, error)

	inventoryv1.UnimplementedInventoryServiceServer
}
