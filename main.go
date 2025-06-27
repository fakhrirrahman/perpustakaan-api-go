package main

import (
	"go-web-native/dto"
	"go-web-native/internal/api"
	"go-web-native/internal/config"
	"go-web-native/internal/connection"
	"go-web-native/internal/repository"
	"go-web-native/internal/seeder"
	"go-web-native/internal/service"
	"log"
	"net/http"
	"time"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "go-web-native/docs" // Import docs for swagger
)

// @title Perpustakaan API
// @version 1.0
// @description API untuk sistem manajemen perpustakaan
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token for authentication

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	gormDB := connection.GetGormDatabase(cnf.Database) // Tambah GORM connection
	
	// Run admin seeder
	log.Println("🌱 Running admin seeder...")
	adminSeeder := seeder.NewAdminSeeder(gormDB)
	if err := adminSeeder.Seed(); err != nil {
		log.Printf("⚠️ Admin seeder failed (this is normal if admin already exists): %v", err)
	}
	
	app := fiber.New(fiber.Config{
		AppName: "Perpustakaan API v1.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// API Info route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Perpustakaan API v1.0",
			"status":  "running",
			"docs":    "/swagger/index.html",
			"endpoints": fiber.Map{
				"authors":     "/api/authors",
				"books":       "/api/books", 
				"customers":   "/api/customers",
				"book_stocks": "/api/book-stocks",
				"auth":        "/api/auth/login",
				"health":      "/api/health",
			},
		})
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Perpustakaan API v1.0",
			"version": "1.0",
			"status":  "healthy",
			"endpoints": fiber.Map{
				"authors":     "/api/authors",
				"books":       "/api/books",
				"customers":   "/api/customers",
				"book_stocks": "/api/book-stocks",
				"auth":        "/api/auth/login",
				"docs":        "/swagger/index.html",
				"health":      "/api/health",
			},
		})
	})

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "healthy",
			"database":  "connected",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	jwtMid := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponeError("endpoint perlu token, silahkan login"))
		},
	})

	// Repositories
	customerRepository := repository.NewCustomerRepository(dbConnection)
	UserRepository := repository.NewUser(gormDB) // Use GORM for User
	bookRepository := repository.NewBook(dbConnection)
	BookStockRepository := repository.NewStock(dbConnection)
	
	// Author repository (GORM)
	authorRepository := repository.NewAuthorRepository(gormDB)

	// Existing services
	customerService := service.NewCustomerService(customerRepository)
	authService := service.NewAuthService(cnf, UserRepository)
	bookService := service.NewBookService(bookRepository, BookStockRepository)
	bookStockService := service.NewBookStock(bookRepository, BookStockRepository)
	
	// Author service
	authorService := service.NewAuthorService(authorRepository)

	// Existing API routes
	api.NewCustomerAPI(app, customerService, jwtMid)
	api.NewAuth(app, authService)
	api.NewBook(app, bookService, jwtMid)
	api.NewBookStock(app, bookStockService, jwtMid)
	
	// Author API routes
	api.NewAuthorAPI(app, authorService, jwtMid)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")	

}
