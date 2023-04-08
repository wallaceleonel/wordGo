package main

import (
	"time"
)

func toDo() {
	type Task struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		DueDate   time.Time `json:"due_date"`
		Completed bool      `json:"completed"`
	}
}
