package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ridhompra/crud-fiber/controllers/bookcontrollers"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect .ENV")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	api := app.Group("/api")

	// book session
	book := api.Group("/books")
	book.Get("/", bookcontrollers.Index)
	book.Get("/:id", bookcontrollers.GetBook)
	book.Post("/", bookcontrollers.CreateBook)
	book.Put("/:id", bookcontrollers.PutBook)
	book.Delete("/:id", bookcontrollers.DeleteBook)
	app.Listen(port)

}
