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
		print("running")

		index, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		UpdateTask("hello.json", index, parts[2])

	}

}
