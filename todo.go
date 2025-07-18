package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	name   string
	status bool
}

// Commands: add-task, remove task, view tasks, complete task,

func main() {

	var file *os.File
	var err error
	var tasks map[string]interface{}

	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		// File doesn't exist — create empty map and file
		tasks = make(map[string]interface{})

		file, err = os.Create("data.json") // creates with write mode
		if err != nil {
			panic(err)
		}
		defer file.Close()

		json.NewEncoder(file).Encode(tasks)

	} else {
		// File exists — open for read
		readFile, err := os.Open("data.json")
		if err != nil {
			panic(err)
		}
		defer readFile.Close()

		// Decode into tasks
		if err := json.NewDecoder(readFile).Decode(&tasks); err != nil {
			panic(err)
		}
	}
	// Open the same file again for write/truncate (for saveFile)
	file, err = os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer saveFile(tasks, file)

	args := os.Args // get cmd line args
	cmd := args[1]

	switch cmd {
	case "add":
		name := args[2]
		addTask(tasks, name)

	case "start":
		name := args[2]
		startTask(tasks, name)

	case "remove":
		name := args[2]
		delete(tasks, name)
		fmt.Println("Task removed Successfully \n ", tasks)

	case "view":
		fmt.Println(tasks)

	case "complete":
		name := args[2]
		tasks[name] = "Completed"

	default:
		fmt.Println("Not a valid command")

	}
}

func addTask(tasks map[string]interface{}, name string) {
	tasks[name] = "Not started"
	fmt.Println("Task added Successfully \n ", tasks)
}

func saveFile(tasks map[string]interface{}, fp *os.File) {
	json.NewEncoder(fp).Encode(tasks)
}

func startTask(tasks map[string]interface{}, task string) {
	tasks[task] = "In progress"
	fmt.Println("Task started Successfully \n ", tasks)
}
