package main

import (
	"fmt"
	"log"
	"os"

	"github.com/macrespo42/Tasker/internal/tasks"
)

func main() {
	userCommand, err := parseUserInput(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("action: %v, id: %v, param: %v\n", userCommand.action, userCommand.id, userCommand.param)
	if userCommand.action == "add" {
		tasks.AddTask(userCommand.param)
	}
}
