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

type BookAPI struct {
	bookService domain.BookService
}

func NewBook(app *fiber.App, bookService domain.BookService, authMidd fiber.Handler) {
	bookAPI := &BookAPI{
		bookService: bookService,
	}
	app.Get("/books", authMidd,bookAPI.Index)
	app.Get("/books/:id", authMidd, bookAPI.Show)
	app.Post("/books", authMidd, bookAPI.Create)
	app.Put("/books/:id", authMidd, bookAPI.Update)
	app.Delete("/books/:id", authMidd, bookAPI.Delete)
}

func (b BookAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := b.bookService.Index(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponeSuccess(res))
}
func (b BookAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponeErrorData("gagal validasi", fails))
	}

	err := b.bookService.Create(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.CreateResponeSuccess(""))
}

func (b BookAPI) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := b.bookService.Show(c, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponeSuccess(res))
}

func (b BookAPI) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	id := ctx.Params("id")
	req.Id = id
	err := b.bookService.Update(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponeSuccess(""))
}

func (b BookAPI) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := b.bookService.Delete(c, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponeSuccess(""))
}


