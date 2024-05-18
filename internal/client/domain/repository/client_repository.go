package repository

import (
	"context"
	"github.com/augustomcosta/client-api/internal/client/domain/entity"
)

type ClientRepository interface {
	Save(ctx context.Context, client *entity.Client) error
}
