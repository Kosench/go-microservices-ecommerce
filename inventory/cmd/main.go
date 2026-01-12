package main

import (
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/part"
	invService "github.com/Kosench/go-microservices-ecommerce/inventory/internal/service/part"
)

func main() {
	repo := part.NewMemoryInventoryRepository()

	_ = invService.NewInventoryService(repo)
}
