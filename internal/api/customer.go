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

func NewCustomerAPI(app *fiber.App, customerService domain.CustomerService) {
	ca := &CustomerAPI{CustomerService: customerService}
	app.Get("/customers", ca.Index) // Perbaikan: Dipindahkan ke dalam fungsi
}

func (ca *CustomerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.CustomerService.Index(c) // Perbaikan: Gunakan `c` langsung
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponeSuccess(res))
}
