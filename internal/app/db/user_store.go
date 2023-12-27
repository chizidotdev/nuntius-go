package db

import (
	"context"
	"github.com/chizidotdev/nuntius/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ domain.UserRepository = (*UserStore)(nil)

type UserStore struct {
	DB *gorm.DB
}

type User struct {
	Base
	FirstName     string
	LastName      string
	Username      string
	Image         string
	Email         string `gorm:"not null;uniqueIndex"`
	EmailVerified bool   `gorm:"not null;default:false"`
	GoogleID      string
	Messages      []Message
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		DB: db,
	}
}

func (s *UserStore) Upsert(_ context.Context, arg *domain.User) error {
	u := User{
		FirstName:     arg.FirstName,
		LastName:      arg.LastName,
		Username:      arg.Username,
		Email:         arg.Email,
		Image:         arg.Image,
		EmailVerified: arg.EmailVerified,
		GoogleID:      arg.GoogleID,
	}
	err := s.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		UpdateAll: true,
	}).Create(&u).Error
	return err
}

func (s *UserStore) Update(_ context.Context, arg *domain.User) error {
	u := User{
		FirstName:     arg.FirstName,
		LastName:      arg.LastName,
		Username:      arg.Username,
		Email:         arg.Email,
		Image:         arg.Image,
		EmailVerified: arg.EmailVerified,
		GoogleID:      arg.GoogleID,
	}
	err := s.DB.Model(&u).Updates(&u).Error
	return err
}

func (s *UserStore) Get(_ context.Context, id uuid.UUID) (*domain.User, error) {
	var u domain.User
	err := s.DB.Preload("Messages").First(&u, "id = ?", id).Error
	return &u, err
}

func (s *UserStore) GetByUsername(_ context.Context, username string) (*domain.User, error) {
	var u domain.User
	err := s.DB.Preload("Messages").First(&u, "username = ?", username).Error
	return &u, err
}

func (s *UserStore) Delete(_ context.Context, id uuid.UUID) error {
	err := s.DB.Delete(&User{}, "id = ?", id).Error
	return err
}
