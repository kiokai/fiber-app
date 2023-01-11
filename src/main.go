package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kiokai/fiber-app/src/routes"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	log.Fatal(app.Listen(":5000"))
}
