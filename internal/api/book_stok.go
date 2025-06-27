package api

import (
	"context"
	"go-web-native/domain"
	"go-web-native/dto"
	"go-web-native/internal/util"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BookStokApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App, bookStockService domain.BookStockService, authMidd fiber.Handler) {
	bookStokApi := &BookStokApi{
		bookStockService: bookStockService,
	}
	app.Post("/api/book-stocks", authMidd, bookStokApi.Create)
	app.Delete("/api/book-stocks", authMidd, bookStokApi.Delete)
}

// Create godoc
// @Summary Create book stock entries
// @Description Create multiple book stock entries for a book
// @Tags book-stocks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book_stock body dto.CreateBookStokData true "Book stock data"
// @Success 201 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 422 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /book-stocks [post]
func (ba *BookStokApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()


	var req dto.CreateBookStokData
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponeErrorData)
	}
	err := ba.bookStockService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeErrorData)
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponeSuccess(""))
	
}
// Delete godoc
// @Summary Delete book stock entries
// @Description Delete multiple book stock entries by codes
// @Tags book-stocks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param code query string true "Semicolon separated codes" example:"BOOK001;BOOK002;BOOK003"
// @Success 200 {object} dto.ResponeString
// @Failure 400 {object} dto.ResponeString
// @Failure 500 {object} dto.ResponeString
// @Router /book-stocks [delete]
func (ba *BookStokApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	codeStr := ctx.Query("code")
	if codeStr == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("parameter code wajib diisi"))
	}

	codes := strings.Split(codeStr,";")


	err := ba.bookStockService.Delete(c, dto.DeleteBookStokData{
		Codes: codes,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponeSuccess("berhasil menghapus stok buku"))
}
	