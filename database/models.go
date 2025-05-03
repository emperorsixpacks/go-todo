package database

import (
	"errors"
	"fmt"
)

func removeAt[T any](s []T, index int) []T {
	if index < 0 || index > len(s) {
		return s
	}
	s = append(s[:index], s[index+1:]...)
	return s

}

func reOrderTasks(tasks []*Task) []*Task {
	for i := 0; i < len(tasks); i++ {
		tasks[i].index = i
	}
	return tasks
}

var DB = GetCache()

type Task struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Completed bool   `json:"completed"`
	index     int
}

type TasksList struct {
	Tasks []*Task `json:"tasks"`
}

type ResponseMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func GetTasks() (TasksList, bool) {
	items, _ := DB.Get("tasks")
	tasks_items, _ := items.(TasksList)
	if len(tasks_items.Tasks) == 0 {
		return TasksList{}, false
	}
	return tasks_items, true
}

func GetTaskbyID(id int64) (*Task, error) {
	var task *Task
	items, ok := GetTasks()
	if !ok {
		return nil, errors.New("No task found")
	}
	for _, x := range items.Tasks {
		if x.ID != id {
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
	task.ID = int64(new_task_id)
	task.index = len(tasks.Tasks)
	tasks.Tasks = append(tasks.Tasks, &task)
	DB.Set("tasks", tasks, 0)
	return nil
}

func DeleteTask(id int64) (TasksList, error) {
	task, err := GetTaskbyID(id)
	if err != nil {
		return TasksList{}, err
	}
	tasks, _ := GetTasks()
	new_tasks := removeAt(tasks.Tasks, task.index)
	tasks.Tasks = reOrderTasks(new_tasks)
	DB.Set("tasks", tasks, 0)
	return tasks, nil
}

func MarkAsComplete(id int64) error {
	task, err := GetTaskbyID(id)
	if err != nil {
		return err
	}
	task.Completed = true
	tasks, _ := GetTasks()
	tasks.Tasks[task.index] = task
	DB.Set("tasks", tasks, 0)
	return nil
}
