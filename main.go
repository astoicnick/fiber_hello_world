package main

import "github.com/gofiber/fiber"

func main(){
	app := fiber.New()

	app.Get("/", indexHandler)

	app.Listen(1010)
}

func indexHandler(c *fiber.Ctx) {
	c.Send("Hello from fiber!")
}