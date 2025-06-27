package seeder

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Seeder interface for all seeders
type Seeder interface {
	Seed() error
}

// DatabaseSeeder manages all database seeders
type DatabaseSeeder struct {
	db      *gorm.DB
	seeders []Seeder
}

// NewDatabaseSeeder creates a new DatabaseSeeder instance
func NewDatabaseSeeder(db *gorm.DB) *DatabaseSeeder {
	return &DatabaseSeeder{
		db:      db,
		seeders: make([]Seeder, 0),
	}
}

// AddSeeder adds a seeder to the list
func (ds *DatabaseSeeder) AddSeeder(seeder Seeder) {
	ds.seeders = append(ds.seeders, seeder)
}

// SeedAll runs all registered seeders
func (ds *DatabaseSeeder) SeedAll() error {
	log.Println("🌱 Starting database seeding...")
	
	for i, seeder := range ds.seeders {
		log.Printf("📦 Running seeder %d/%d...", i+1, len(ds.seeders))
		
		if err := seeder.Seed(); err != nil {
			return fmt.Errorf("seeder %d failed: %w", i+1, err)
		}
	}
	
	log.Printf("🎉 All %d seeders completed successfully!", len(ds.seeders))
	return nil
}

// SetupDefaultSeeders sets up the default seeders
func (ds *DatabaseSeeder) SetupDefaultSeeders() {
	// Add admin seeder
	adminSeeder := NewAdminSeeder(ds.db)
	ds.AddSeeder(adminSeeder)
	
	// Future: Add more seeders here
	// bookSeeder := NewBookSeeder(ds.db)
	// ds.AddSeeder(bookSeeder)
	
	// customerSeeder := NewCustomerSeeder(ds.db)
	// ds.AddSeeder(customerSeeder)
}
