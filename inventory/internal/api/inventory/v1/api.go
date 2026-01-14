package v1

import (
	"context"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/converter"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/service"
	inventoryv1 "github.com/Kosench/go-microservices-ecommerce/shared/pkg/proto/inventory/v1"
)

type api struct {
	inventoryv1.UnimplementedInventoryServiceServer

	invService service.Service
}

func NewAPI(invService service.Service) *api {
	return &api{
		invService: nil,
	}
}

func (a *api) GetPart(ctx context.Context, req *inventoryv1.GetPartRequest) (*inventoryv1.GetPartResponse, error) {
	part, err := a.invService.GetPart(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	grpcPart := converter.ConvertPartToGRPC(part)
	return &inventoryv1.GetPartResponse{
		Part: grpcPart,
	}, nil
}

func (a *api) ListParts(ctx context.Context, req *inventoryv1.ListPartsRequest) (*inventoryv1.ListPartsResponse, error) {
	filter := converter.ConvertFilterFromGRPC(req.GetFilter())
	parts, err := a.invService.ListParts(ctx, filter)
	if err != nil {
		return nil, err
	}

	grpcParts := make([]*inventoryv1.Part, len(parts))
	for i, part := range parts {
		grpcParts[i] = converter.ConvertPartToGRPC(part)
	}

	return &inventoryv1.ListPartsResponse{
		Parts: grpcParts,
	}, nil
}
