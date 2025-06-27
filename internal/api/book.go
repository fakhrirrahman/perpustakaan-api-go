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
	app.Get("/api/books", authMidd, bookAPI.Index)
	app.Get("/api/books/:id", authMidd, bookAPI.Show)
	app.Post("/api/books", authMidd, bookAPI.Create)
	app.Put("/api/books/:id", authMidd, bookAPI.Update)
	app.Delete("/api/books/:id", authMidd, bookAPI.Delete)
}

// Index godoc
// @Summary Get all books
// @Description Get list of all books
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ResponeArrayBookData
// @Failure 500 {object} dto.ResponeString
// @Router /books [get]
func (b BookAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := b.bookService.Index(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponeSuccess(res))
}
// Create godoc
// @Summary Create new book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body dto.CreateBookRequest true "Book data"
// @Success 201 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 422 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /books [post]
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

// Show godoc
// @Summary Get book by ID
// @Description Get a single book by ID
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} dto.ResponeBookData
// @Failure 400 {object} dto.ResponeString
// @Failure 404 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /books/{id} [get]
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

// Update godoc
// @Summary Update book
// @Description Update an existing book
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Param book body dto.UpdateBookRequest true "Book data"
// @Success 200 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 404 {object} dto.ResponeString
// @Failure 422 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /books/{id} [put]
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

// Delete godoc
// @Summary Delete book
// @Description Delete a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 404 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /books/{id} [delete]
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


