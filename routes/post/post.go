package post

import (
	"time"

	"github.com/furkanalptekin/go-crud/database"
	"github.com/furkanalptekin/go-crud/models"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/posts/:id", getById)
	app.Get("/posts", get)
	app.Post("/posts", post)
	app.Put("/posts/:id", put)
	app.Delete("/posts/:id", delete)
}

func getById(c *fiber.Ctx) error {
	post := models.Post{}
	id := c.Params("id")
	database.Instance.Find(&post, id)

	if post.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func get(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.Instance.Find(&posts)
	return c.Status(fiber.StatusOK).JSON(posts)
}

func post(c *fiber.Ctx) error {
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx := database.Instance.Create(&post)
	if tx.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": tx.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(post)
}

func put(c *fiber.Ctx) error {
	post := models.Post{}
	id := c.Params("id")
	database.Instance.Find(&post, id)

	if post.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	request := new(models.Post)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx := database.Instance.Model(&post).UpdateColumns(map[string]interface{}{
		"title":      request.Title,
		"body":       request.Body,
		"updated_at": time.Now(),
	})

	if tx.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": tx.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func delete(c *fiber.Ctx) error {
	post := models.Post{}
	id := c.Params("id")
	database.Instance.Find(&post, id)

	if post.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	tx := database.Instance.Table("posts").Where("id = ?", id).Delete(&models.Post{})
	if tx.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": tx.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
