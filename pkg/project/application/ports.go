package application

import (
	"context"

	"github.com/techforge-lat/bastion/pkg/project/domain"
	"github.com/techforge-lat/dafi/v2"
)

type Repository interface {
	Create(ctx context.Context, req domain.CreateProjectRequest) error
	UpdateByCriteria(ctx context.Context, req domain.UpdateProjectRequest, criteria dafi.Criteria) error
	DeleteByCriteria(ctx context.Context, criteria dafi.Criteria) error
	GetByCriteria(ctx context.Context, criteria dafi.Criteria) (domain.Project, error)
	ListByCriteria(ctx context.Context, criteria dafi.Criteria) (domain.Projects, error)
}
