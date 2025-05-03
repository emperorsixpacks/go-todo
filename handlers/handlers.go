package handlers

import (
	"fmt"

	databse "github.com/emperorsixpacks/go-task/database"
	"github.com/gofiber/fiber/v2"
)

var DB = databse.GetCache()

type Task struct {
	Id           int    `json:"page-id"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Is_completed bool   `json:"is-completed"`
}

type TasksList struct {
	Tasks []Task `json:"tasks"`
}

type Message struct {
	ErrorMessage string `json:"message"`
	StatusCode   int    `json:"status"`
}

func getTasks() ([]TasksList, bool) {
	items, ok := DB.Get("tasks")
	if !ok {
		return []TasksList{}, false
	}
	return items.([]TasksList), true
}

func geTask(id int)

func GetTasks(ctx *fiber.Ctx) error {
	items, ok := getTasks()
	if !ok {
		message := Message{"No tasks found now", fiber.StatusNotFound}
		return ctx.Status(message.StatusCode).JSON(message)
	}

	return ctx.JSON(items)
}

func GetTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	items, ok := getTasks()
	if !ok {
		message_str := fmt.Sprintf("Task with id %s not found", id)
		message := Message{message_str, fiber.StatusNotFound}
		return ctx.Status(message.StatusCode).JSON(message)
	}

	return ctx.JSON(items)
}

func DeleteTask(ctx *fiber.Ctx) error {
	return nil
}

func UpdateTask(ctx *fiber.Ctx) error {
	return nil
}

func CreateTask(ctx *fiber.Ctx) error {
	return nil
}
