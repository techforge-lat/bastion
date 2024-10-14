package domain

import (
	"errors"

	"github.com/techforge-lat/bastion/pkg/kit/domain"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/errortrace/v2/errtype"
)

type DisplayName string

func (d DisplayName) Validate() error {
	if len(d) == 0 {
		return errortrace.Wrap(errors.New("DisplayName is required")).SetErrCode(errtype.UnprocessableEntity)
	}

	if len(d) > 100 {
		return errortrace.Wrap(errors.New("DisplayName is too long")).SetErrCode(errtype.UnprocessableEntity)
	}

	return nil
}

type Description string

type Project struct {
	domain.Base
	DisplayName DisplayName `json:"display_name"`
	Description Description `json:"description"`
}

type Projects []Project

type CreateProjectRequest struct {
	ID          domain.ID        `json:"id"`
	DisplayName DisplayName      `json:"display_name"`
	Description Description      `json:"description"`
	CreatedAt   domain.CreatedAt `json:"created_at"`
}

func (c CreateProjectRequest) Validate() error {
	if err := c.ID.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	if err := c.DisplayName.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	if err := c.CreatedAt.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	return nil
}

type UpdateProjectRequest struct {
	DisplayName DisplayName      `json:"display_name"`
	Description Description      `json:"description"`
	UpdatedAt   domain.UpdatedAt `json:"created_at"`
}

func (c UpdateProjectRequest) Validate() error {
	if err := c.DisplayName.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	if err := c.UpdatedAt.Validate(); err != nil {
		return errortrace.Wrap(err)
	}

	return nil
}
