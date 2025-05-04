# 📝 Go TODO API

A simple task manager API built with [Fiber](https://github.com/gofiber/fiber) in Go, using an in-memory cache for storing tasks.

---

## 🚀 Features

- List all tasks
- View a single task
- Create a task
- Delete a task
- Mark a task as complete
- In-memory task caching using Go's `cache`

---

## 📦 Project Structure

```

go-todo/
├── database/       # In-memory database (cache) and task model
├── handlers/       # Request handlers for API endpoints
├── main.go         # App entrypoint

````

---

## 📋 API Endpoints

| Method | Endpoint         | Description              |
|--------|------------------|--------------------------|
| GET    | `/todos`         | Get all tasks            |
| GET    | `/todos/:id`     | Get task by ID           |
| POST   | `/create`        | Create a new task        |
| PUT    | `/complete/:id`  | Mark task as completed   |
| DELETE | `/delete/:id`    | Delete a task by ID      |

---

## 🛠️ Setup

### Requirements

- Go 1.18 or later

### Installation

```bash
git clone https://github.com/emperorsixpacks/go-todo.git
cd go-todo
go mod tidy
go run main.go
````

The server will start on: [http://localhost:3000](http://localhost:3000)

---

## 📂 Sample Task Structure

```json
{
  "id": "1",
  "title": "Learn Go Fiber",
  "completed": false
}
```

---

## 🧠 Notes

* All data is stored in-memory using a cache (no persistent database)-ttl 2mins.
* Ideal for learning Go and building basic REST APIs with Fiber.

---

## 🧑‍💻 Author

[emperorsixpacks](https://github.com/emperorsixpacks)

---

## 📄 License

MIT License

