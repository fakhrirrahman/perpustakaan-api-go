package api

import (
	"context"
	"go-web-native/domain"
	"go-web-native/dto"
	"go-web-native/internal/util"
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
	app.Post("/customers", ca.Create)
	app.Put("/customers/:id", ca.Update)
}

func (ca CustomerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.CustomerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponeSuccess(res))
}

func (ca CustomerAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponeErrorData("validation error", fails))
}
err := ca.CustomerService.Create(c, req)
if err != nil {
	return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
}
return ctx.Status(http.StatusCreated).JSON(dto.CreateResponeSuccess(""))
}

func (ca CustomerAPI) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponeErrorData("validation error", fails))
	}
	req.ID = ctx.Params("id")
	err := ca.CustomerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponeSuccess(""))
}