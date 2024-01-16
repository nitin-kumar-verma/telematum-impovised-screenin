package handler

import (
	"fmt"
	"screening/internal/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	if name == "" || email == "" {
		//More validations can be added here
		return c.Status(fiber.StatusBadRequest).SendString("Name and Email can't be empty")
	}
	id, err := db.InsertUser(name, email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Inserted user with id %d", id))
}

func UpdateUser(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	id := c.FormValue("id")

	if name == "" || email == "" || id == "" {
		//More validations can be added here
		return c.Status(fiber.StatusBadRequest).SendString("Name, Email and Id can't be empty")
	}
	idVal, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Id should be int")
	}
	updateCount, err := db.UpdateUser(name, email, idVal)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Updated %d users Successfully", updateCount))
}
