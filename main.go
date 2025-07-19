package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	userCommand, err := parseUserInput(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("action: %v, id: %v, param: %v", userCommand.action, userCommand.id, userCommand.param)
}
