package service

import (
	"context"
	"sync"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/converter"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/model"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository"
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
	repoPart, err := s.partRepository.GetPart(ctx, uuid)
	if err != nil {
		return nil, err
	}

	servicePart := converter.ConvertRepoPartToModelPart(repoPart)
	return servicePart, nil
}

func (s *inventoryService) ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error) {
	repoFilter := converter.ConvertModelPartsFilterToRepoPartsFilter(filter)
	repoParts, err := s.partRepository.ListParts(ctx, repoFilter)
	if err != nil {
		return nil, err
	}

	serviceParts := converter.ConvertRepoPartsToModelParts(repoParts)
	return serviceParts, nil
}
