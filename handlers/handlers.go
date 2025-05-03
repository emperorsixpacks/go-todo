package handlers

import (
	"encoding/json"
	"log"

	databse "github.com/emperorsixpacks/go-todo/database"
	"github.com/gofiber/fiber/v2"
)

var DB = databse.GetCache()

type Todo struct {
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Is_completed bool   `json:"is-completed"`
}

type Page struct {
	Id    int    `json:"page-id"`
	Todos []Todo `json:"todos"`
}

func GetTodos(ctx *fiber.Ctx) error {
	item, ok := DB.Get("todos")
	if !ok {
		return ctx.Status(fiber.StatusNotFound).SendString("Not todos found at the momenet, try creating some now ")
	}
	item_json, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("internal server err")
	}
	return ctx.JSON(item_json)
}

func GetTodo(ctx *fiber.Ctx) error {
	return nil
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
