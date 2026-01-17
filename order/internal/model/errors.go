package model

import (
	"errors"

	sharedErrors "github.com/Kosench/go-microservices-ecommerce/shared/pkg/errors"
)

var (
	ErrOrderNotFound          = sharedErrors.NewNotFoundError(errors.New("order not found"))
	ErrOrderAlreadyPaid       = sharedErrors.NewInvalidArgumentError(errors.New("order already paid"))
	ErrOrderCannotBeCancelled = sharedErrors.NewInvalidArgumentError(errors.New("order cannot be cancelled"))
	ErrInvalidOrderUUID       = sharedErrors.NewInvalidArgumentError(errors.New("invalid order uuid"))
)
