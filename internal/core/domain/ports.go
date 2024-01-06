package domain

import (
	"context"
	"github.com/google/uuid"
)

type (
	UserStore interface {
		Upsert(context.Context, *User) (*User, error)
		UpdateUsername(context.Context, *User) (*User, error)
		Get(context.Context, uuid.UUID) (*User, error)
		GetByUsername(context.Context, string) (*User, error)
		Delete(context.Context, uuid.UUID) error
	}

	MessageStore interface {
		Create(context.Context, *Message) error
		Get(context.Context, uuid.UUID) (*Message, error)
		GetAll(context.Context, uuid.UUID) ([]*Message, error)
		Delete(context.Context, uuid.UUID) error
		UpdateArchive(context.Context, UpdateArchiveInput) error
	}

	UpdateArchiveInput struct {
		ID       uuid.UUID
		Archived bool
	}
)
