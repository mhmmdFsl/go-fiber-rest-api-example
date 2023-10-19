package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Fiber Rest Example",
	})
	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") == "" {
			return fiber.NewError(400, "Path variable cant be empty")
		}
		return c.SendString("Hello " + c.Params("name") + " !")
	})

	app.Listen(":3030")
}
