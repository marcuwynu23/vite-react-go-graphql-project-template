package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/example/fullstack-template/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	u, err := s.users.FindByEmail(ctx, strings.TrimSpace(in.Email))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(in.Password) == "" {
		return nil, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return &LoginResult{Token: mustToken()}, nil
}

func (s *AuthService) SeedInitialUser(ctx context.Context, email, password string) error {
	if s.users == nil || !s.users.Enabled() {
		return nil
	}
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if email == "" || password == "" {
		return nil
	}

	_, err := s.users.FindByEmail(ctx, email)
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.users.Create(ctx, email, string(hash))
}

func mustToken() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "demo-token"
	}
	return hex.EncodeToString(b)
}
