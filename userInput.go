package main

import (
	"errors"
	"strconv"
)

type command struct {
	action string
	id     int
	param  string
}

func parseUserInput(args []string) (command, error) {
	if len(args) < 2 {
		return command{}, errors.New("Please provide an action")
	}

	userInput := command{action: args[1]}
	if len(args) > 2 {
		if id, err := strconv.Atoi(args[2]); err == nil {
			userInput.id = id
		} else {
			userInput.param = args[2]
		}
	}

	if len(args) == 4 {
		userInput.param = args[3]
	}

	return userInput, nil
}
