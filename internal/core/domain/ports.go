package domain

import (
	"context"
	"github.com/google/uuid"
)

type (
	UserRepository interface {
		Upsert(context.Context, *User) error
		Update(context.Context, *User) error
		Get(context.Context, uuid.UUID) (*User, error)
		GetByUsername(context.Context, string) (*User, error)
		Delete(context.Context, uuid.UUID) error
	}

	MessageRepository interface {
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
