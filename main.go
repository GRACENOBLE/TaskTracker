package main

import (
	// "time"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Todo CLI: ")

		command, _ := reader.ReadString('\n')

		parts := strings.Fields(command)

		if len(parts) < 2 {
			fmt.Println("-> Please input more than one argument.")
			continue
		}

		if parts[0] == "list" && parts[1] != "" {
			data := ReadFile("todos.json")

			if parts[1] == "all" {
				for _, todo := range data {

					fmt.Printf("%v\n", todo)

				}
				continue
			} else if parts[1] == "done" {
				for _, todo := range data {
					if todo.Status == Done {
						fmt.Printf("%v\n", todo)
					}
				}
			} else if parts[1] == "todo" {
				for _, todo := range data {
					if todo.Status == Todo {
						fmt.Printf("%v\n", todo)
					}
				}
			} else if parts[1] == "in-progress" {
				for _, todo := range data {
					if todo.Status == InProgress {
						fmt.Printf("%v\n", todo)
					}
				}
			}
		} else if parts[0] == "add" && parts[1] != "" {

			AppendToFile("todos.json", parts[1])

		} else if parts[0] == "update" && parts[1] != "" && parts[2] != "" {

			if len(parts) == 3 {

				index, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				//remember to give warning in case user accidentally creates a todo with name same as status enum type

				UpdateTask("todos.json", index, parts[2], "")

			} else if len(parts) == 4 {

				index, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				UpdateTask("todos.json", index, parts[2], parts[3])

			}

		} else if parts[0] == "mark" && parts[1] != "" && parts[2] != "" {

			if len(parts) == 3 {

				index, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				UpdateTask("todos.json", index, "", parts[2])

			}

		} else if parts[0] == "delete" && parts[1] != "" {

			fmt.Println("-> Are you sure?(y/n)")

			var choice string
			fmt.Scanf("%v", &choice)

			if choice == "Y" || choice == "y" {

				index, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				DeleteTask("todos.json", index)

			}

		} else {
			fmt.Println("Oops! I dont seem to know that command. Try the \"help\" command")
		}

	}
}
