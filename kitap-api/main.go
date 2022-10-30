package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{Title: "LOTR", Author: "J.R.R TOLİEN"},
	{Title: "Fahrenheit451", Author: "Ray Bradbury"},
	{Title: "Dune", Author: "Frank Herbert"},
}

func main() {

	app := fiber.New()

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	app.Listen(":5055")
}
