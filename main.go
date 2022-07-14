package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "Jose",
		Lastname:  "Dominguez",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.NewString()

	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	//Middleware
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about!")
	})
	/*
		app.Get("/user", handleUser)
		app.Post("/users", handleCreateUser)
	*/

	app.Use(requestid.New())
	userGroup := app.Group("/users")

	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}
