package service

import (
	"context"
	"sync"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/model"
)

type inventoryService struct {
	mu             sync.RWMutex
	partRepository repository.InventoryRepository
}

func NewInventoryService(inventoryRepository repository.InventoryRepository) *inventoryService {
	return &inventoryService{
		partRepository: inventoryRepository,
	}
}

func (s *inventoryService) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	part, err := s.partRepository.GetPart(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *inventoryService) ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error) {
	parts, err := s.partRepository.ListParts(ctx, filter)
	if err != nil {
		return nil, err
	}

	return parts, nil
}
