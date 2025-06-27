package seeder

import (
	"fmt"
	"go-web-native/domain"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AdminSeeder handles the seeding of admin user
type AdminSeeder struct {
	db *gorm.DB
}

// NewAdminSeeder creates a new AdminSeeder instance
func NewAdminSeeder(db *gorm.DB) *AdminSeeder {
	return &AdminSeeder{db: db}
}

// Seed creates the default admin user
func (s *AdminSeeder) Seed() error {
	// Check if admin already exists
	var existingUser domain.User
	result := s.db.Where("email = ?", "admin@perpustakaan.com").First(&existingUser)
	
	if result.Error == nil {
		log.Println("🔍 Admin user already exists, skipping seeder...")
		return nil
	}

	// Hash the default password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create admin user
	adminUser := domain.User{
		ID:        uuid.New().String(),
		Email:     "admin@perpustakaan.com",
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert admin user
	if err := s.db.Create(&adminUser).Error; err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	log.Println("✅ Admin user created successfully!")
	log.Println("📧 Email: admin@perpustakaan.com")
	log.Println("🔑 Password: admin123")
	log.Println("⚠️  Please change the default password in production!")

	return nil
}

// SeedMultipleAdmins creates multiple admin users for testing
func (s *AdminSeeder) SeedMultipleAdmins() error {
	admins := []struct {
		email    string
		password string
	}{
		{"admin@perpustakaan.com", "admin123"},
		{"superadmin@perpustakaan.com", "super123"},
		{"librarian@perpustakaan.com", "librarian123"},
	}

	for _, admin := range admins {
		// Check if user already exists
		var existingUser domain.User
		result := s.db.Where("email = ?", admin.email).First(&existingUser)
		
		if result.Error == nil {
			log.Printf("🔍 User %s already exists, skipping...", admin.email)
			continue
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("❌ Failed to hash password for %s: %v", admin.email, err)
			continue
		}

		// Create user
		user := domain.User{
			ID:        uuid.New().String(),
			Email:     admin.email,
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := s.db.Create(&user).Error; err != nil {
			log.Printf("❌ Failed to create user %s: %v", admin.email, err)
			continue
		}

		log.Printf("✅ User %s created successfully!", admin.email)
	}

	return nil
}
