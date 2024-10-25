package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateFile() {
	todo := Task{
		Id: 1, Description: "Hello",
		Status:    InProgress,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	todoJson, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error converting Struct to JSON: %v", err)
	}
	e:= os.WriteFile("Hello.json", todoJson, 0644)
	if e != nil {
		log.Fatalf("Error converting Struct to JSON: %v", e)
		}
	fmt.Println("File Written successfully")

	fmt.Printf("%v", todo)

}

func ReadFile() {

	content, err := os.ReadFile("Hello.json")
	var data Task

	if err != nil {
		log.Fatalf("Error reading file %v", err)
	}

	e := json.Unmarshal(content, &data)
	if e != nil {
		log.Fatalf("Error unmarshalling the data: %v", e)
	}
	fmt.Printf("%v", data)

}
