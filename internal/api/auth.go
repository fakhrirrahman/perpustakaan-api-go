package api

import (
	"context"
	"go-web-native/domain"
	"go-web-native/dto"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthApi struct {
	AuthService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	aa := AuthApi{
		AuthService: authService,
	}

	app.Post("/api/auth/login", aa.Login)
}

// Login godoc
// @Summary User login
// @Description Authenticate user and get access token
// @Tags authentication
// @Accept json
// @Produce json
// @Param credentials body dto.AuthData true "Login credentials"
// @Success 200 {object} dto.ResponeAuthResponse
// @Failure 400 {object} dto.ResponeString
// @Failure 401 {object} dto.ResponeString
// @Failure 422 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /auth/login [post]
func (aa AuthApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()
	var req dto.AuthData
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, err := aa.AuthService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponeSuccess(res))
}