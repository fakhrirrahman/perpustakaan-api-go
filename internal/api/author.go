package api

import (
	"go-web-native/domain"
	"go-web-native/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthorHandler struct {
	authorService domain.AuthorService
}

func NewAuthorHandler(authorService domain.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Get list of all authors
// @Tags authors
// @Accept json
// @Produce json
// @Success 200 {object} dto.Respone[[]dto.AuthorData]
// @Failure 500 {object} dto.Respone[string]
// @Router /authors [get]
func (h *AuthorHandler) GetAllAuthors(c *fiber.Ctx) error {
	authors, err := h.authorService.Index(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}

	response := dto.Respone[[]dto.AuthorData]{
		Code:    http.StatusOK,
		Message: "Berhasil mengambil data penulis",
		Data:    authors,
	}

	return c.Status(http.StatusOK).JSON(response)
}

// GetAuthorById godoc
// @Summary Get author by ID
// @Description Get a single author by ID
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} dto.Respone[dto.AuthorData]
// @Failure 400 {object} dto.Respone[string]
// @Failure 404 {object} dto.Respone[string]
// @Failure 500 {object} dto.Respone[string]
// @Router /authors/{id} [get]
func (h *AuthorHandler) GetAuthorById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("ID penulis diperlukan"))
	}

	author, err := h.authorService.Show(c.Context(), id)
	if err != nil {
		if err == domain.AuthorNotFound {
			return c.Status(http.StatusNotFound).JSON(dto.CreateResponeError(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}

	response := dto.Respone[dto.AuthorData]{
		Code:    http.StatusOK,
		Message: "Berhasil mengambil data penulis",
		Data:    author,
	}

	return c.Status(http.StatusOK).JSON(response)
}

// CreateAuthor godoc
// @Summary Create new author
// @Description Create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body dto.CreateAuthorRequest true "Author data"
// @Success 201 {object} dto.Respone[string]
// @Failure 400 {object} dto.Respone[string]
// @Failure 409 {object} dto.Respone[string]
// @Failure 500 {object} dto.Respone[string]
// @Security BearerAuth
// @Router /authors [post]
func (h *AuthorHandler) CreateAuthor(c *fiber.Ctx) error {
	var req dto.CreateAuthorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("Format data tidak valid"))
	}

	err := h.authorService.Create(c.Context(), req)
	if err != nil {
		if err == domain.AuthorEmailAlreadyExists {
			return c.Status(http.StatusConflict).JSON(dto.CreateResponeError(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}

	response := dto.Respone[string]{
		Code:    http.StatusCreated,
		Message: "Berhasil membuat penulis baru",
		Data:    "",
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// UpdateAuthor godoc
// @Summary Update author
// @Description Update an existing author
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Param author body dto.UpdateAuthorRequest true "Author data"
// @Success 200 {object} dto.Respone[string]
// @Failure 400 {object} dto.Respone[string]
// @Failure 404 {object} dto.Respone[string]
// @Failure 409 {object} dto.Respone[string]
// @Failure 500 {object} dto.Respone[string]
// @Security BearerAuth
// @Router /authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("ID penulis diperlukan"))
	}

	var req dto.UpdateAuthorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("Format data tidak valid"))
	}

	// Set ID dari parameter URL
	req.ID = id

	err := h.authorService.Update(c.Context(), req)
	if err != nil {
		if err == domain.AuthorNotFound {
			return c.Status(http.StatusNotFound).JSON(dto.CreateResponeError(err.Error()))
		}
		if err == domain.AuthorEmailAlreadyExists {
			return c.Status(http.StatusConflict).JSON(dto.CreateResponeError(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}

	response := dto.Respone[string]{
		Code:    http.StatusOK,
		Message: "Berhasil mengupdate penulis",
		Data:    "",
	}

	return c.Status(http.StatusOK).JSON(response)
}

// DeleteAuthor godoc
// @Summary Delete author
// @Description Delete an author by ID
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} dto.Respone[string]
// @Failure 400 {object} dto.Respone[string]
// @Failure 404 {object} dto.Respone[string]
// @Failure 500 {object} dto.Respone[string]
// @Security BearerAuth
// @Router /authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponeError("ID penulis diperlukan"))
	}

	err := h.authorService.Delete(c.Context(), id)
	if err != nil {
		if err == domain.AuthorNotFound {
			return c.Status(http.StatusNotFound).JSON(dto.CreateResponeError(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponeError(err.Error()))
	}

	response := dto.Respone[string]{
		Code:    http.StatusOK,
		Message: "Berhasil menghapus penulis",
		Data:    "",
	}

	return c.Status(http.StatusOK).JSON(response)
}

// NewAuthorAPI untuk setup routing Author
func NewAuthorAPI(app *fiber.App, authorService domain.AuthorService, jwtMid fiber.Handler) {
	handler := NewAuthorHandler(authorService)

	// Author routes
	authorGroup := app.Group("/api/authors")
	
	// Public routes (bisa diakses tanpa JWT)
	authorGroup.Get("/", handler.GetAllAuthors)
	authorGroup.Get("/:id", handler.GetAuthorById)
	
	// Protected routes (perlu JWT)
	authorGroup.Post("/", jwtMid, handler.CreateAuthor)
	authorGroup.Put("/:id", jwtMid, handler.UpdateAuthor) 
	authorGroup.Delete("/:id", jwtMid, handler.DeleteAuthor)
}
