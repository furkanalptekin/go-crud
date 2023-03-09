package main

import (
	"github.com/furkanalptekin/go-crud/database"
	"github.com/furkanalptekin/go-crud/routes/post"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()
	post.Init(app)

	app.Listen(":3000")
}
