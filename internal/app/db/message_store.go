package db

import (
	"context"
	"github.com/chizidotdev/nuntius/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

var _ domain.MessageStore = (*MessageStore)(nil)

type MessageStore struct {
	DB *gorm.DB
}

type Message struct {
	Base
	Content  string
	Archived bool
	UserID   uuid.UUID
}

func NewMessageStore(db *gorm.DB) *MessageStore {
	err := db.AutoMigrate(&Message{})
	if err != nil {
		log.Fatal("Failed to migrate Message")
	}

	return &MessageStore{
		DB: db,
	}
}

func (m *MessageStore) Create(ctx context.Context, arg *domain.Message) error {
	msg := Message{
		Content:  arg.Content,
		Archived: arg.Archived,
		UserID:   arg.UserID,
	}
	err := m.DB.WithContext(ctx).Create(&msg).Error
	return err
}

func (m *MessageStore) Get(ctx context.Context, id uuid.UUID) (*domain.Message, error) {
	msg := Message{}
	err := m.DB.WithContext(ctx).Where("id = ?", id).First(&msg).Error
	return &domain.Message{
		ID:       msg.ID,
		Content:  msg.Content,
		Archived: msg.Archived,
		UserID:   msg.UserID,
	}, err
}

func (m *MessageStore) GetAll(ctx context.Context, userID uuid.UUID) ([]*domain.Message, error) {
	var msgs []*domain.Message
	err := m.DB.WithContext(ctx).
		Order("created_at desc").
		Find(&msgs, "user_id = ?", userID).
		Error
	return msgs, err
}

func (m *MessageStore) Delete(ctx context.Context, id uuid.UUID) error {
	err := m.DB.WithContext(ctx).Delete(&Message{}, "id = ?", id).Error
	return err
}

func (m *MessageStore) UpdateArchive(ctx context.Context, arg domain.UpdateArchiveInput) error {
	err := m.DB.WithContext(ctx).
		Model(&Message{}).
		Where("id = ?", arg.ID).
		Update("archived", arg.Archived).Error
	return err
}
