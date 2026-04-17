package repository

import (
	"context"

	"github.com/example/fullstack-template/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Enabled() bool {
	return r != nil && r.db != nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	if r.db == nil {
		return nil, gorm.ErrRecordNotFound
	}
	var u models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(ctx context.Context, email, passwordHash string) error {
	if r.db == nil {
		return gorm.ErrInvalidDB
	}
	u := &models.User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	return r.db.WithContext(ctx).Create(u).Error
}
