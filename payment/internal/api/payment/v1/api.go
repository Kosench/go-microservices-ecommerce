package v1

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/payment/internal/convert"
	"github.com/Kosench/go-microservices-ecommerce/payment/internal/service"
	paymentV1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/payment/v1"
)

type api struct {
	paymentV1.UnimplementedPaymentServiceServer

	payService service.Service
}

func NewAPI(srv service.Service) *api {
	return &api{
		payService: srv,
	}
}
func (a *api) PayOrder(ctx context.Context, req *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	payment := convert.ConvertFromGRPC(req)

	transactionUUID, err := a.payService.PayOrder(ctx, payment)
	if err != nil {
		return nil, err
	}

	return &paymentV1.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}
