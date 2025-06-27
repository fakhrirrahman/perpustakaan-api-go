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

func NewCustomerAPI(app *fiber.App, customerService domain.CustomerService, AuzMidd fiber.Handler) {
	ca := CustomerAPI{
		CustomerService: customerService,
	}
	app.Get("/api/customers", AuzMidd, ca.Index)
	app.Post("/api/customers", AuzMidd, ca.Create)
	app.Put("/api/customers/:id", AuzMidd, ca.Update)
	app.Delete("/api/customers/:id", AuzMidd, ca.Delete)
	app.Get("/api/customers/:id", AuzMidd, ca.Show)
}

// Index godoc
// @Summary Get all customers
// @Description Get list of all customers
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ResponeArrayCustomerData
// @Failure 500 {object} dto.ResponeString
// @Router /customers [get]
func (ca CustomerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.CustomerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponeSuccess(res))
}

// Create godoc
// @Summary Create new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param customer body dto.CreateCustomerRequest true "Customer data"
// @Success 201 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 422 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /customers [post]
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

// Update godoc
// @Summary Update customer
// @Description Update an existing customer
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Param customer body dto.UpdateCustomerRequest true "Customer data"
// @Success 200 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 404 {object} dto.ResponeString
// @Failure 422 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /customers/{id} [put]
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

// Delete godoc
// @Summary Delete customer
// @Description Delete a customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ResponeString
// @Failure 404 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /customers/{id} [delete]
func (ca CustomerAPI) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ca.CustomerService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.SendStatus(http.StatusNoContent)
}

// Show godoc
// @Summary Get customer by ID
// @Description Get a single customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} dto.ResponeCustomerData
// @Failure 400 {object} dto.ResponeString
// @Failure 404 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /customers/{id} [get]
func (ca CustomerAPI) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := ca.CustomerService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponeSuccess(res))
}