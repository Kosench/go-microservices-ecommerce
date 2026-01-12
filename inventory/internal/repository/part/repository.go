package part

import (
	"sync"

	inventoryv1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/inventory/v1"
)

type repository struct {
	mu   sync.RWMutex
	data map[string]*inventoryv1.Part
}

func NewInventoryRepository() *repository {
	return &repository{
		data: make(map[string]*inventoryv1.Part),
	}
}
