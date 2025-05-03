package main

import (
	"log"

	"github.com/emperorsixpacks/go-todo/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
  setUpRoutes(app)
	log.Printf("Starting service on port :3000")
	app.Listen(":3000")
}

func setUpRoutes(a *fiber.App) {
	a.Get("/todos", handlers.GetTodos)
	a.Get("/todos/:id", handlers.GetTodo)
	a.Get("/delete/:id", handlers.DeleteTodo)
	a.Get("/update/:id", handlers.UpdateTodo)
	a.Get("/create", handlers.CreateTodo)
}
