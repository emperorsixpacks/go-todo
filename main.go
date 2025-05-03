package main

import (
	"log"

	database "github.com/emperorsixpacks/go-todo/database"
	"github.com/emperorsixpacks/go-todo/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	initDB()
	setUpRoutes(app)
	log.Printf("Starting service on port :3000")
	app.Listen(":3000")
}

func setUpRoutes(a *fiber.App) {
	a.Get("/todos", handlers.GetTasks)
	a.Get("/todos/:id", handlers.GetTask)
	a.Delete("/delete/:id", handlers.DeleteTask)
	a.Put("/complete/:id", handlers.UpdateTask)
	a.Post("/create", handlers.CreateTask)
}

func initDB() {
	_tasksList := database.TasksList{Tasks: make([]*database.Task, 0)}
	db := database.GetCache()
	db.Set("tasks", _tasksList, 0)
}
