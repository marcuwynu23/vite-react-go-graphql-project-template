package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/example/fullstack-template/internal/repository"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	users *repository.UserRepository
}

func NewAuthService(users *repository.UserRepository) *AuthService {
	return &AuthService{users: users}
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
}

func (s *AuthService) Login(ctx context.Context, in LoginInput) (*LoginResult, error) {
	if s.users == nil || !s.users.Enabled() {
		return &LoginResult{Token: mustToken()}, nil
	}
	_, err := s.users.FindByEmail(ctx, in.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}
	// Password verification would go here (e.g. bcrypt.CompareHashAndPassword).
	if in.Password == "" {
		return nil, ErrInvalidCredentials
	}
	return &LoginResult{Token: mustToken()}, nil
}

func mustToken() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "demo-token"
	}
	return hex.EncodeToString(b)
}
