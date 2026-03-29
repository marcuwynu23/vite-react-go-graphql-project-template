package handler

import (
	"errors"

	"github.com/example/fullstack-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body service.LoginInput
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid json")
	}
	res, err := h.auth.Login(c.Context(), body)
	if errors.Is(err, service.ErrInvalidCredentials) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(res)
}
