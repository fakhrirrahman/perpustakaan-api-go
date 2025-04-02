package api

import (
	"context"
	"go-web-native/domain"
	"go-web-native/dto"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CustomerAPI struct {
	CustomerService domain.CustomerService
}

func NewCustomerAPI(app *fiber.App, customerService domain.CustomerService){
	ca := CustomerAPI{
		CustomerService: customerService,
	}
	app.Get("/customers", ca.Index)
}

func (ca CustomerAPI) Index(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ca.CustomerService.Index(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return c.JSON(dto.CreateResponeSuccess(res))
}
