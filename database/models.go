package database

import (
	"errors"
	"fmt"
)

var DB Cache

type Task struct {
	id           int64
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Is_completed bool   `json:"is-completed"`
}

type TasksList struct {
	Tasks []*Task `json:"tasks"`
}

type ResponseMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func GetTasks() (TasksList, bool) {
	items, ok := DB.Get("tasks")
	if !ok {
		return TasksList{}, false
	}
	return items.(TasksList), true
}

func GetTaskbyID(id int64) (*Task, error) {
	var task *Task
	items, ok := GetTasks()
	if !ok {
		return nil, errors.New("No task found")
	}
	for _, x := range items.Tasks {
		if x.id != id {
			continue
		}
		task = x
	}
	if task == nil {
		return nil, fmt.Errorf("No task with id %d", id)
	}
	return task, nil
}

func CreateTask(task Task) error {
	tasks, _ := GetTasks()
	new_task_id := len(tasks.Tasks) + 1
	task.id = int64(new_task_id)
	tasks.Tasks = append(tasks.Tasks, &task)
	DB.Set("todos", tasks, 0)
	return nil
}

func deleteTask(id int64) error {
	_, err := GetTaskbyID(id)
	if err != nil {
		return err
	}
	return nil
}
