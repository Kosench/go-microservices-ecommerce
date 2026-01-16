package service

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/payment/internal/model"
)

type Service interface {
	PayOrder(ctx context.Context, req model.Pay) (string, error)
}
