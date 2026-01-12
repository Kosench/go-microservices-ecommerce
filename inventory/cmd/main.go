package main

import (
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/part"
	service "github.com/Kosench/go-microservices-ecommerce/inventory/internal/service/part"
)

func main() {
	repo := part.NewMemoryInventoryRepository()

	svc := service.NewInventoryService(repo)
}
