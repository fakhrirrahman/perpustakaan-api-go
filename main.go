package main

import (
	"go-web-native/internal/config"
	"go-web-native/internal/connection"
	"go-web-native/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	CustomerRepository := repository.NewCustomerRepository(dbConnection)

	app.Get("/developers", developers)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")	

}
