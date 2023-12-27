package domain

import "github.com/google/uuid"

type (
	User struct {
		ID            uuid.UUID
		FirstName     string
		LastName      string
		Username      string
		Email         string
		Image         string
		EmailVerified bool
		GoogleID      string
		Messages      []Message
	}

	Message struct {
		ID       uuid.UUID
		Content  string
		Archived bool
		UserID   uuid.UUID
	}
)
