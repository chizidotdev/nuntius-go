package service

import (
	"context"
	"encoding/gob"
	"github.com/chizidotdev/nuntius/internal/core/domain"
	"github.com/google/uuid"
)

type MessageService struct {
	repo     domain.MessageStore
	userRepo domain.UserStore
}

func NewMessageService(repo domain.MessageStore, userRepo domain.UserStore) *MessageService {
	gob.Register(domain.Message{})

	return &MessageService{
		repo:     repo,
		userRepo: userRepo,
	}
}

type CreateMessageReq struct {
	Content  string `form:"content" binding:"required"`
	Username string
}

func (m *MessageService) CreateMessage(ctx context.Context, req CreateMessageReq) error {
	user, err := m.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return err
	}

	return m.repo.Create(ctx, &domain.Message{
		Content:  req.Content,
		Archived: false,
		UserID:   user.ID,
	})
}

func (m *MessageService) ListMessages(ctx context.Context, userID uuid.UUID) ([]*domain.Message, error) {
	return m.repo.GetAll(ctx, userID)
}
