package repository

import (
	"context"

	"github.com/alochym01/web-w-golang/project/models"
)

// PostRepo with all its methods
type PostRepo interface {
	Fetch(ctx context.Context, numberRecord int64) ([]*models.Post, error)
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	Create(ctx context.Context, p *models.Post) (int64, error)
	Update(ctx context.Context, p *models.Post) (*models.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
