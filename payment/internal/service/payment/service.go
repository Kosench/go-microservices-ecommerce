package payment

import (
	"context"
	"log"

	"github.com/Kosench/go-microservices-ecommerce/payment/internal/model"
	"github.com/Kosench/go-microservices-ecommerce/payment/internal/service"
	"github.com/google/uuid"
)

type payService struct{}

func NewPayService() service.Service {
	return &payService{}
}

func (s *payService) PayOrder(ctx context.Context, req model.Pay) (string, error) {
	transactionUUID := uuid.New().String()

	// Имитируем обработку платежа
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", transactionUUID)

	return transactionUUID, nil
}
