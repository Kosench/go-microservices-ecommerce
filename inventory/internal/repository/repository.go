package repository

import "context"

type Repository interface {
	GetPart(ctx context.Context, uuid string)
}
