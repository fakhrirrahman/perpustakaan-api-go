package main

import (
	"go-web-native/internal/api"
	"go-web-native/internal/config"
	"go-web-native/internal/connection"
	"go-web-native/internal/repository"
	"go-web-native/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	customerRepository := repository.NewCustomerRepository(dbConnection)

	customerService := service.NewCustomerService(customerRepository)

	api.NewCustomerAPI(app, customerService)


	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")	

}
