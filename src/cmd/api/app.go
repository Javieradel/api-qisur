package main

import (
	"log"

	"github.com/Javieradel/api-qisur.git/src/db"
	"github.com/gofiber/fiber/v3"
)

func main() {
	db.InitDB()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("HELLO")
	})

	log.Fatal(app.Listen(":3000"))
}
