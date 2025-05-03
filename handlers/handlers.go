package handlers

import (
	"fmt"
	"log"
	"strconv"

	database "github.com/emperorsixpacks/go-todo/database"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(ctx *fiber.Ctx) error {
	items, ok := database.GetTasks()
	if !ok {
		message := database.Message{ErrorMessage: "No tasks found now", StatusCode: fiber.StatusNotFound}
		return ctx.Status(message.StatusCode).JSON(message)
	}

	return ctx.JSON(items)
}

func GetTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	int_id, _ := strconv.ParseInt(id, 10, 64)
	item, err := database.GetTaskbyID(int_id)
	if err != nil {
		log.Fatal(err)
		message_str := fmt.Sprintf("Task with id %s not found", id)
		message := database.Message{ErrorMessage: message_str, StatusCode: fiber.StatusNotFound}
		return ctx.Status(message.StatusCode).JSON(message)
	}

	return ctx.JSON(item)
}

func DeleteTask(ctx *fiber.Ctx) error {
	return nil
}

func UpdateTask(ctx *fiber.Ctx) error {
	return nil
}

func CreateTask(ctx *fiber.Ctx) error {
	var task database.Task
	if err := ctx.BodyParser(&task); err != nil {
		message := database.Message{ErrorMessage: "Unprocessable request", StatusCode: fiber.StatusUnprocessableEntity}
		return ctx.Status(message.StatusCode).JSON(message)
	}
}
