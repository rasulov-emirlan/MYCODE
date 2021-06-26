package main

import "github.com/gofiber/fiber/v2"

func helloPage(c *fiber.Ctx) error {
	return c.SendString(("jdidmmadkmclsmdlk"))
}

func main() {
	app := fiber.New()

	app.Get("/", helloPage)
	app.Listen(":8080")
}
