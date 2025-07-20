package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func getTasks() ([]Task, error) {
	file, err := os.ReadFile("./db/tasks.json")

	if err != nil {
		log.Fatal(err)
	}

	if len(file) == 0 {
		return make([]Task, 0), nil
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return make([]Task, 0), err
	}
	return tasks, nil
}

func AddTask(description string) {
	tasks, err := getTasks()
	if err != nil {
		log.Fatal(err)
	}
	newTask := Task{Id: len(tasks) + 1, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	tasks = append(tasks, newTask)

	fmt.Println("____________________DEBUG_________________")
	fmt.Println(newTask)
	fmt.Println(tasks)
	fmt.Println("____________________DEBUG_________________")

	file, err := os.OpenFile("./db/tasks.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(tasks)
}
