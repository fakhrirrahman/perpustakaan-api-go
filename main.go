package main

import (
	"go-web-native/dto"
	"go-web-native/internal/api"
	"go-web-native/internal/config"
	"go-web-native/internal/connection"
	"go-web-native/internal/repository"
	"go-web-native/internal/service"
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	jwtMid := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponeError("endpoint perlu token, silahkan login"))
		},
	})

	customerRepository := repository.NewCustomerRepository(dbConnection)
	UserRepository := repository.NewUser(dbConnection)


	customerService := service.NewCustomerService(customerRepository)
	authService := service.NewAuthService(cnf, UserRepository)
	

	api.NewCustomerAPI(app, customerService, jwtMid)
	api.NewAuth(app, authService)



	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")	

}
