package main

import (

	//"log"
	"time"
	//"fmt"
)

func main() {
	todo := Task{
		Id: 1, Description: "Hello",
		Status:    InProgress,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	AppendToFile("hello.json", todo)
	// data := ReadFile("hello.json")
	// fmt.Print(data)

}
