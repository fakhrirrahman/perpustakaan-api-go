package connection

import (
	"database/sql"
	"fmt"
	"go-web-native/domain"
	"go-web-native/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(conf config.Database) *sql.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        conf.User,
        conf.Pass,
        conf.Host,
        conf.Port,
        conf.Name,
    )
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    return db
}

// GetGormDatabase untuk GORM connection
func GetGormDatabase(conf config.Database) *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        conf.User,
        conf.Pass,
        conf.Host,
        conf.Port,
        conf.Name,
    )
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error opening GORM database: %v", err)
    }

    // Auto migrate semua tables
    err = db.AutoMigrate(
        &domain.Author{},
        &domain.Book{},
        &domain.Customer{},
        &domain.User{},
        &domain.BookStock{},
    )
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

    log.Println("✅ Database migration completed successfully!")
    return db
}
