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
	app.Post("/book-stocks", authMidd, bookStokApi.Create)
	app.Delete("/book-stocks", authMidd, bookStokApi.Delete)
}

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
func (ba *BookStokApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	codes := strings.Split(ctx.Query("code"),";")
	if len(codes) < 1 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("parameter code wajib diisi"))
	}

	err := ba.bookStockService.Delete(c, dto.DeleteBookStokData{
		Codes: codes,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponeSuccess("berhasil menghapus stok buku"))
}
	