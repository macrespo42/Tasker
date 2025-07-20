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

	switch userCommand.action {
	case "add":
		tasks.Add(userCommand.param)
	case "list":
		tasks.List(userCommand.param)
	case "update":
		tasks.Update(userCommand.id, userCommand.param)
	case "mark-in-progress":
		tasks.UpdateStatus(userCommand.id, "in-progress")
	case "mark-done":
		tasks.UpdateStatus(userCommand.id, "done")
	case "delete":
		tasks.Delete(userCommand.id)
	default:
		panic("unrecognized command")
	}
}
