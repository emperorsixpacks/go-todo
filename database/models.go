package database

import (
	"errors"
	"fmt"
)

var DB Cache

type Task struct {
	Id           int64  `json:"page-id"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Is_completed bool   `json:"is-completed"`
}

type TasksList struct {
	Tasks []*Task `json:"tasks"`
}

type Message struct {
	ErrorMessage string `json:"message"`
	StatusCode   int    `json:"status"`
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
		if x.Id != id {
			continue
		}
		task = x
	}
	if task == nil {
		return nil, fmt.Errorf("No task with id %d", id)
	}
	return task, nil
}

func deleteTask(id int64) error {
	_, err := GetTaskbyID(id)
	if err != nil {
		return err
	}
	return nil
}
