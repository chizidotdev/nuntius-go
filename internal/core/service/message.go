package service

import (
	"encoding/gob"
	"github.com/chizidotdev/nuntius/internal/core/domain"
)

type MessageService struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) *MessageService {
	gob.Register(domain.Message{})

	return &MessageService{
		repo: repo,
	}
}
