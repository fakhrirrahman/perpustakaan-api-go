package domain

import (
	"context"
	"go-web-native/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthData) (dto.AuthResponse, error)
}