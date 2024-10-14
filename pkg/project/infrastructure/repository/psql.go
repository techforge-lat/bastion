package repository

import (
	"context"

	"github.com/techforge-lat/bastion/internal/database"
	"github.com/techforge-lat/bastion/pkg/project/domain"
	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type PsqlRepositoryAdapter struct {
	db *database.Adapter
}

func NewPsqlRepositoryAdapter(db *database.Adapter) *PsqlRepositoryAdapter {
	return &PsqlRepositoryAdapter{
		db: db,
	}
}

func (p *PsqlRepositoryAdapter) Create(ctx context.Context, req domain.CreateProjectRequest) error {
	result, err := createQuery.ToSql()
	if err != nil {
		return errortrace.Wrap(err)
	}

	if _, err := p.db.Exec(ctx, result.Sql, req.ID, req.DisplayName, req.Description, req.CreatedAt); err != nil {
		return errortrace.Wrap(err)
	}

	return nil
}

func (p *PsqlRepositoryAdapter) UpdateByCriteria(ctx context.Context, req domain.UpdateProjectRequest, criteria dafi.Criteria) error {
	result, err := updateQuery.WithValues(req.DisplayName, req.Description, req.UpdatedAt).Where(criteria.Filters...).ToSQL()
	if err != nil {
		return errortrace.Wrap(err)
	}

	if _, err := p.db.Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.Wrap(err)
	}

	return nil
}

func (p *PsqlRepositoryAdapter) DeleteByCriteria(ctx context.Context, criteria dafi.Criteria) error {
	result, err := deleteQuery.Where(criteria.Filters...).ToSQL()
	if err != nil {
		return errortrace.Wrap(err)
	}

	if _, err := p.db.Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.Wrap(err)
	}

	return nil
}

func (p *PsqlRepositoryAdapter) GetByCriteria(ctx context.Context, criteria dafi.Criteria) (domain.Project, error) {
	result, err := getQuery.Where(criteria.Filters...).OrderBy(criteria.Sorts...).Limit(criteria.Pagination.PageSize).Page(criteria.Pagination.PageNumber).ToSQL()
	if err != nil {
		return domain.Project{}, errortrace.Wrap(err)
	}

	m := domain.Project{}
	if err := p.db.QueryRow(ctx, result.Sql).Scan(&m.ID, &m.DisplayName, &m.Description, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return domain.Project{}, errortrace.Wrap(err)
	}

	return m, nil
}

func (p *PsqlRepositoryAdapter) ListByCriteria(ctx context.Context, criteria dafi.Criteria) (domain.Projects, error) {
	panic("not implemented") // TODO: Implement
}
