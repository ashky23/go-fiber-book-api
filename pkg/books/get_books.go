package books

import (
	"github.com/ashky23/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&books)
}
