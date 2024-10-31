package main

import (
	// "time"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	print("What would you like to do?\n")

	command, _ := reader.ReadString('\n')

	parts := strings.Fields(command)

	if parts[0] == "add" && parts[1] != "" {

		AppendToFile("hello.json", parts[1])

	} else if parts[0] == "update" && parts[1] != "" && parts[2] != "" {
		if len(parts) == 3 {
			index, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			//remember to give warning in case user accidentally creates a todo with name same as ststus enum type

			UpdateTask("hello.json", index, parts[2], "")
		} else if len(parts) == 4 {

			index, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			UpdateTask("hello.json", index, parts[2], parts[3])
		}
	} else if parts[0] == "mark" && parts[1] != "" && parts[2] != "" {

		if len(parts) == 3 {

			index, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			UpdateTask("hello.json", index, "", parts[2])

		}

	}

}
