package repository

import (
	"context"
	"go-web-native/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) FindByEmail(ctx context.Context, email string) (usr domain.User, err error) {
	err = u.db.WithContext(ctx).Where("email = ?", email).First(&usr).Error
	if err != nil {
		return domain.User{}, err
	}
	return usr, nil
}	