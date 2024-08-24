package main

import (
	"fmt"
	"os"

	"task-cli/task"
)

func main() {
	filename := "./tasks.json"
	tasks, err := task.LoadTask(filename)

	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Expected a command")
		return
	}

	command := args[0]

	switch command {

	case "add":
		if len(args) < 2 {
			fmt.Println("Expected description for the task")
			return
		}
		task.AddTask(&tasks, args[1])

	case "update":
		if len(args) < 3 {
			fmt.Println("Expected task ID and new description")
			return
		}
		id := task.ParseID(args[1])
		task.UpdateTask(&tasks, id, args[2])

	case "delete":
		if len(args) < 2 {
			fmt.Println("Expected task ID")
			return
		}
		id := task.ParseID(args[1])
		task.DeleteTask(&tasks, id)

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Expected task ID")
			return
		}
		id := task.ParseID(args[1])
		task.MarkTask(&tasks, id, task.StatusInProgress)

	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Expected task ID")
			return
		}
		id := task.ParseID(args[1])
		task.MarkTask(&tasks, id, task.StatusDone)

	case "list":
		status := ""

		if len(args) > 1 {
			status = args[1]
		}
		task.ListTasks(tasks, status)

	default:
		fmt.Println("Unknown command:", command)

	}

	if err := task.SaveTask(filename, tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	}

}
