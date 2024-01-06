package db

import (
	"context"
	"github.com/chizidotdev/nuntius/internal/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

var _ domain.UserStore = (*UserStore)(nil)

type UserStore struct {
	DB *gorm.DB
}

type User struct {
	Base
	FirstName     string
	LastName      string
	Username      string `gorm:"uniqueIndex"`
	Image         string
	Email         string `gorm:"not null;unique"`
	EmailVerified bool   `gorm:"not null;default:false"`
	GoogleID      string
	Messages      []Message
}

func NewUserStore(db *gorm.DB) *UserStore {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate Message")
	}

	return &UserStore{
		DB: db,
	}
}

func (s *UserStore) Upsert(ctx context.Context, arg *domain.User) (*domain.User, error) {
	u := User{
		FirstName:     arg.FirstName,
		LastName:      arg.LastName,
		Username:      arg.Username,
		Email:         arg.Email,
		Image:         arg.Image,
		EmailVerified: arg.EmailVerified,
		GoogleID:      arg.GoogleID,
	}
	err := s.DB.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "email"}},
			DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "image", "email_verified", "google_id"}),
		}, clause.Returning{}).
		Create(&u).Error

	return &domain.User{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Username:      u.Username,
		Email:         u.Email,
		Image:         u.Image,
		EmailVerified: u.EmailVerified,
		GoogleID:      u.GoogleID,
	}, err
}

func (s *UserStore) UpdateUsername(ctx context.Context, arg *domain.User) (*domain.User, error) {
	u := User{}
	err := s.DB.WithContext(ctx).
		Model(&u).
		Clauses(clause.Returning{}).
		Where("email = ?", arg.Email).
		Update("username", arg.Username).Error

	return &domain.User{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Username:      u.Username,
		Email:         u.Email,
		Image:         u.Image,
		EmailVerified: u.EmailVerified,
		GoogleID:      u.GoogleID,
	}, err
}

func (s *UserStore) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var u domain.User
	err := s.DB.WithContext(ctx).Preload("Messages").First(&u, "id = ?", id).Error
	return &u, err
}

func (s *UserStore) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var u domain.User
	err := s.DB.WithContext(ctx).Preload("Messages").First(&u, "username = ?", username).Error
	return &u, err
}

func (s *UserStore) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.DB.WithContext(ctx).Delete(&User{}, "id = ?", id).Error
	return err
}
