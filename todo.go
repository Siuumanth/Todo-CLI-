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
	defer handlePanic()

	args := os.Args // get cmd line args

	if len(args) < 2 {
		panic("Please enter a command")
	}

	cmd := args[1]

	switch cmd {
	case "add":
		verifyArgs(len(args), "Task name should be entered")
		name := args[2]
		addTask(tasks, name)

	case "start":
		name := args[2]
		startTask(tasks, name)

	case "remove":
		verifyArgs(len(args), "Task name should be entered")
		name := args[2]
		delete(tasks, name)
		fmt.Println("Task removed Successfully \n ")
		showTasks(tasks)

	case "show":
		showTasks(tasks)

	case "complete":
		verifyArgs(len(args), "Task name should be entered")
		name := args[2]
		tasks[name] = "Completed"
		fmt.Println("Task completed Successfully \n ")
		showTasks(tasks)

	case "help":
		fmt.Println(`
		Available commands:
		add <task name> - Add a new task
		start <task name> - Start a task
		remove <task name> - Remove a task
		show - show all tasks
		complete <task name> - Mark a task complete
		help - Show this message
		`)

	default:
		fmt.Println("Not a valid command")

	}
}

func verifyArgs(args int, msg string) {
	if args < 3 {
		panic(msg)
	}
	if args > 3 {
		panic("Too many arguments")
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}

func addTask(tasks map[string]interface{}, name string) {
	tasks[name] = "Not started"
	fmt.Println("Task added Successfully \n ")
	showTasks(tasks)

}

func saveFile(tasks map[string]interface{}, fp *os.File) {
	json.NewEncoder(fp).Encode(tasks)
}

func startTask(tasks map[string]interface{}, task string) {
	tasks[task] = "In progress"
	fmt.Println("Task started Successfully \n ")
	showTasks(tasks)

}

func showTasks(tasks map[string]interface{}) {
	table := [][]string{{"Task", "Status"}}
	for task, status := range tasks {
		table = append(table, []string{task, fmt.Sprint(status)})
	}

	fmt.Println(renderTable(table))
}

func renderTable(data [][]string) string {
	table := "  +---------------------------------------------+\n"
	for _, row := range data {
		table += fmt.Sprintf("  | %-20s | %-20s |\n", row[0], row[1])
	}
	table += "  +---------------------------------------------+\n"
	return table
}
