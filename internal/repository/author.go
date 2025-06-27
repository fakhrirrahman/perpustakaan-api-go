package repository

import (
	"context"
	"go-web-native/domain"

	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) domain.AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (r *authorRepository) FindAll(ctx context.Context) ([]domain.Author, error) {
	var authors []domain.Author
	
	if err := r.db.WithContext(ctx).Find(&authors).Error; err != nil {
		return nil, err
	}
	
	return authors, nil
}

func (r *authorRepository) FindById(ctx context.Context, id string) (domain.Author, error) {
	var author domain.Author
	
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return author, domain.AuthorNotFound
		}
		return author, err
	}
	
	return author, nil
}

func (r *authorRepository) Save(ctx context.Context, author *domain.Author) error {
	// Cek apakah email sudah ada
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.Author{}).Where("email = ?", author.Email).Count(&count).Error; err != nil {
		return err
	}
	
	if count > 0 {
		return domain.AuthorEmailAlreadyExists
	}
	
	if err := r.db.WithContext(ctx).Create(author).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *authorRepository) Update(ctx context.Context, author *domain.Author) error {
	// Cek apakah author ada
	if _, err := r.FindById(ctx, author.ID); err != nil {
		return err
	}
	
	// Cek apakah email sudah digunakan oleh author lain
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.Author{}).Where("email = ? AND id != ?", author.Email, author.ID).Count(&count).Error; err != nil {
		return err
	}
	
	if count > 0 {
		return domain.AuthorEmailAlreadyExists
	}
	
	if err := r.db.WithContext(ctx).Save(author).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *authorRepository) Delete(ctx context.Context, id string) error {
	// Cek apakah author ada
	if _, err := r.FindById(ctx, id); err != nil {
		return err
	}
	
	if err := r.db.WithContext(ctx).Delete(&domain.Author{}, "id = ?", id).Error; err != nil {
		return err
	}
	
	return nil
}
