package main

import (
	"log"
	"os"

	"github.com/macrespo42/Tasker/internal/tasks"
)

func main() {
	userCommand, err := parseUserInput(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("action: %v, id: %v, param: %v\n", userCommand.action, userCommand.id, userCommand.param)
	switch userCommand.action {
	case "add":
		tasks.Add(userCommand.param)
	case "list":
		tasks.List(userCommand.param)
	case "delete":
		tasks.Delete(userCommand.id)
	default:
		panic("unrecognized command")
	}
}
