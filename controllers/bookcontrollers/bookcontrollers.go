package bookcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ridhompra/crud-fiber/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"message": "Data tidak ditemukan",
		// })
	}
	return c.JSON(book)
}
func CreateBook(c *fiber.Ctx) error {
	var books models.Book
	if err := c.BodyParser(&books); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := models.DB.Create(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(books)

}
func DeleteBook(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})

	}
	return c.JSON(book)

}
func PutBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if models.DB.Where("id=?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "data berhasil diupdate",
	})
}
