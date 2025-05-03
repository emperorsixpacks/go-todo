package handlers

import (
	"fmt"

	databse "github.com/emperorsixpacks/go-todo/database"
	"github.com/gofiber/fiber/v2"
)

var DB = databse.GetCache()

type Todo struct {
	Id           int    `json:"page-id"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Is_completed bool   `json:"is-completed"`
}

type TodosList struct {
	Todos []Todo `json:"todos"`
}

type Message struct {
	ErrorMessage string `json:"message"`
	StatusCode   int    `json:"status"`
}

func getTodos() ([]TodosList, bool) {
	items, ok := DB.Get("todos")
	if !ok {
		return []TodosList{}, false
	}
	return items.([]TodosList), true
}

func GetTodos(ctx *fiber.Ctx) error {
	items, ok := DB.Get("todos")
	if !ok {
		message := Message{"No todos found now", fiber.StatusNotFound}
		return ctx.Status(message.StatusCode).JSON(message)
	}

	return ctx.JSON(items)
}

func GetTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	items, ok := DB.Get("todos")
	if !ok {
		message_str := fmt.Sprintf("Todo with id %s not found", id)
		message := Message{message_str, fiber.StatusNotFound}
		return ctx.Status(message.StatusCode).JSON(message)
	}

	return ctx.JSON(items)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	return nil
}

func UpdateTodo(ctx *fiber.Ctx) error {
	return nil
}

func CreateTodo(ctx *fiber.Ctx) error {
	return nil
}
