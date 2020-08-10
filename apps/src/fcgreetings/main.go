package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func Greeting(greeting string) string {
	// set default values -- the proper way
	if greeting == "" {
		greeting = "Code.education Rocks!"
	}

	return greeting
}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) {
		// Render index template
		_ = c.Render("index", fiber.Map{
			"Greetings": Greeting(""),
		})
	})

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(3000))
}
