package main

type status string

const (
	Todo       status = "todo"
	InProgress status = "InProgress"
	Done       status = "Done"
)

type Task struct {
	Id          int64
	Description string
	Status      status
	CreatedAt   string
	UpdatedAt   string
}
