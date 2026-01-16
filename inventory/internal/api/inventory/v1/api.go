package v1

import (
	"context"
	"errors"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/converter"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/part"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/service"
	inventoryV1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	inventoryV1.UnimplementedInventoryServiceServer

	invService service.Service
}

func NewAPI(invService service.Service) *api {
	return &api{
		invService: invService,
	}
}

func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.invService.GetPart(ctx, req.GetUuid())
	if err != nil {
		return nil, mapErrorToGRPCStatus(err)
	}

	grpcPart := converter.ConvertPartToGRPC(part)
	return &inventoryV1.GetPartResponse{
		Part: grpcPart,
	}, nil
}

func (a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	filter := converter.ConvertFilterFromGRPC(req.GetFilter())
	parts, err := a.invService.ListParts(ctx, filter)
	if err != nil {
		return nil, mapErrorToGRPCStatus(err)
	}

	grpcParts := make([]*inventoryV1.Part, len(parts))
	for i, part := range parts {
		grpcParts[i] = converter.ConvertPartToGRPC(part)
	}

	return &inventoryV1.ListPartsResponse{
		Parts: grpcParts,
	}, nil
}

// mapErrorToGRPCStatus преобразует внутренние ошибки в gRPC статус коды
func mapErrorToGRPCStatus(err error) error {
	if errors.Is(err, part.ErrNotFound) {
		return status.Error(codes.NotFound, "part not found")
	}
	if errors.Is(err, part.ErrAlreadyExists) {
		return status.Error(codes.AlreadyExists, "part already exists")
	}
	// Для всех остальных ошибок возвращаем Internal
	return status.Error(codes.Internal, "internal server error")
}
