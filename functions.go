package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func IsFileEmpty(filename string) (bool, error) {
	// Get file information
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, err // Return error if file does not exist or any other error occurs
	}

	// Check if file size is 0
	return fileInfo.Size() == 0, nil
}

func ReadFile(file string) []Task {

	var data []Task

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	e := json.Unmarshal(content, &data)
	if e != nil {
		log.Fatalf("Error unmarshalling the data: %v", e)
	}
	return data

}

func AppendToFile(filename string, task string) error {

	var fileData []Task

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}

	empty, err := IsFileEmpty(filename)
	if err != nil {
		log.Fatal(err)
	}

	if empty {
		data := Task{
			Id:          1,
			Description: task,
			Status:      Todo,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		}

		var newArray []Task
		newArray = append(newArray, data)

		todoJson, err := json.MarshalIndent(newArray, "", " ")
		if err != nil {
			log.Fatalf("Error converting Struct to JSON: %v", err)
		}

		e := os.WriteFile("hello.json", todoJson, 0644)
		if e != nil {
			log.Fatalf("Error converting Struct to JSON: %v", e)
		}

		fmt.Printf("File Written successfully")

	} else {

		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error reading file %v", err)
		}
		e := json.Unmarshal(content, &fileData)
		if e != nil {
			log.Fatalf("Error unmarshalling the data: %v", e)
		}

		data := Task{
			Id:          int64(len(fileData) + 1),
			Description: task,
			Status:      Todo,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		}

		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)

		if err != nil {
			return err
		}

		defer file.Close()

		newData := append(fileData, data)

		jsondata, err := json.MarshalIndent(newData, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		if _, err := file.WriteString(string(jsondata) + "\n"); err != nil {
			return err
		}

	}

	return nil

}

func UpdateTask(filename string, id int64, taskBody string, taskStatus string) {

	tasks := ReadFile(filename)
	//will do binary search in future for faster search

	for i, task := range tasks {
		if task.Id == id {
			var updatedTask Task
			if taskStatus == "" {

				updatedTask = Task{
					Id:          task.Id,
					Description: taskBody,
					Status:      task.Status,
					CreatedAt:   task.CreatedAt,
					UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				}

			} else if taskBody == "" {

				updatedTask = Task{
					Id:          task.Id,
					Description: task.Description,
					Status:      status(taskStatus),
					CreatedAt:   task.CreatedAt,
					UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				}

			} else {

				updatedTask = Task{
					Id:          task.Id,
					Description: taskBody,
					Status:      status(taskStatus),
					CreatedAt:   task.CreatedAt,
					UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				}

			}

			tasks[i] = updatedTask
		}
	}

	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	e := os.WriteFile("hello.json", jsonData, 0644)
	if e != nil {
		log.Fatal(e)
	}

}

func DeleteTask(filename string, id int64) {
	tasks := ReadFile(filename)
	//will do binary search in future for faster search
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	jsondata, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	e := os.WriteFile("hello.json", jsondata, 0644)
	if e != nil {
		log.Fatal(e)
	}

	
}
