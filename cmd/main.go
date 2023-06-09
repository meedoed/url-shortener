package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("https://ya.ru", http.StatusSeeOther)
	})

	log.Fatal(app.Listen(":8000"))
}
