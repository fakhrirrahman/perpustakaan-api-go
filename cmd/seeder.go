package main

import (
	"flag"
	"fmt"
	"go-web-native/internal/config"
	"go-web-native/internal/connection"
	"go-web-native/internal/seeder"
	"log"
	"os"
)

func main() {
	var seedType string
	flag.StringVar(&seedType, "type", "admin", "Type of seeder to run (admin, all, multiple)")
	flag.Parse()

	// Load configuration
	cnf := config.Get()
	
	// Get database connection
	gormDB := connection.GetGormDatabase(cnf.Database)
	
	// Create seeder
	adminSeeder := seeder.NewAdminSeeder(gormDB)
	
	switch seedType {
	case "admin":
		log.Println("🌱 Running admin seeder...")
		if err := adminSeeder.Seed(); err != nil {
			log.Fatalf("❌ Admin seeder failed: %v", err)
		}
		
	case "multiple":
		log.Println("🌱 Running multiple admins seeder...")
		if err := adminSeeder.SeedMultipleAdmins(); err != nil {
			log.Fatalf("❌ Multiple admins seeder failed: %v", err)
		}
		
	case "all":
		log.Println("🌱 Running all seeders...")
		dbSeeder := seeder.NewDatabaseSeeder(gormDB)
		dbSeeder.SetupDefaultSeeders()
		if err := dbSeeder.SeedAll(); err != nil {
			log.Fatalf("❌ Database seeder failed: %v", err)
		}
		
	default:
		fmt.Printf("❌ Unknown seeder type: %s\n", seedType)
		fmt.Println("Available types: admin, all, multiple")
		os.Exit(1)
	}
	
	log.Println("🎉 Seeding completed successfully!")
}
