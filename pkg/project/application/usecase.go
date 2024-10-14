package application

import (
	"context"

	"github.com/techforge-lat/bastion/pkg/project/domain"
	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(ctx context.Context, req domain.CreateProjectRequest) error {
	if err := req.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	return errortrace.Wrap(u.repo.Create(ctx, req)).Err()
}

func (u UseCase) UpdateByCriteria(ctx context.Context, req domain.UpdateProjectRequest, criteria dafi.Criteria) error {
	if err := req.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	return errortrace.Wrap(u.repo.UpdateByCriteria(ctx, req, criteria)).Err()
}

func (u UseCase) DeleteByCriteria(ctx context.Context, criteria dafi.Criteria) error {
	return errortrace.Wrap(u.repo.DeleteByCriteria(ctx, criteria)).Err()
}

func (u UseCase) GetByCriteria(ctx context.Context, criteria dafi.Criteria) (domain.Project, error) {
	result, err := u.repo.GetByCriteria(ctx, criteria)
	if err != nil {
		return domain.Project{}, errortrace.Wrap(err)
	}

	return result, nil
}

func (u UseCase) ListByCriteria(ctx context.Context, criteria dafi.Criteria) (domain.Projects, error) {
	result, err := u.repo.ListByCriteria(ctx, criteria)
	if err != nil {
		return nil, errortrace.Wrap(err)
	}

	return result, nil
}
