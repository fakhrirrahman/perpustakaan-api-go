package service

import (
	"context"
	"errors"
	"go-web-native/domain"
	"go-web-native/dto"
	"go-web-native/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	conf *config.Config
	UserRepository domain.UserRepository
}

func NewAuthService(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return &AuthService{
		conf: cnf,
		UserRepository: userRepository,
	}
}

func (a AuthService) Login(ctx context.Context, req dto.AuthData) (dto.AuthResponse, error) {
    user, err := a.UserRepository.FindByEmail(ctx, req.Email)
    if err != nil {
        return dto.AuthResponse{}, err
    }
    if user.ID == "" {
        return dto.AuthResponse{}, errors.New("Authentication failed")
    }
    
    // Verifikasi password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return dto.AuthResponse{}, errors.New("Authentication failed")
    }

    // Generate JWT token
    claim := jwt.MapClaims{
        "id":  user.ID,
        "exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
    }

    // Deklarasikan token dan hasilkan signed string
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // Deklarasikan token di sini
    tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key)) // Tangkap error

    if err != nil {
        return dto.AuthResponse{}, errors.New("Authentication failed")
    }

    return dto.AuthResponse{
        Token: tokenStr,
    }, nil
}
