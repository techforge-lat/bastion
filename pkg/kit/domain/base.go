package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/errortrace/v2/errtype"
	"gopkg.in/guregu/null.v4"
)

type ID uuid.UUID

func (i ID) Validate() error {
	if uuid.UUID(i) == uuid.Nil {
		return errortrace.Wrap(errors.New("ID is required")).SetErrCode(errtype.UnprocessableEntity)
	}

	return nil
}

type CreatedAt time.Time

func (c CreatedAt) Validate() error {
	if time.Time(c).IsZero() {
		return errortrace.Wrap(errors.New("UpdatedAt is required")).SetErrCode(errtype.UnprocessableEntity)
	}

	return nil
}

type UpdatedAt null.Time

func (u UpdatedAt) Validate() error {
	if !u.Valid {
		return errortrace.Wrap(errors.New("UpdatedAt is required")).SetErrCode(errtype.UnprocessableEntity)
	}

	return nil
}

type Base struct {
	ID        ID        `json:"id"`
	CreatedAt CreatedAt `json:"created_at"`
	UpdatedAt UpdatedAt `json:"updated_at"`
}
