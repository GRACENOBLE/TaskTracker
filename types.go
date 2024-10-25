package main

type status int

const (
	Todo       status = iota //0
	InProgress               //1
	Done                     //2
)

type Task struct {
	Id          int64
	Description string
	Status      status
	CreatedAt   string
	UpdatedAt   string
}
