package task

import (
	"fmt"
	"os"
	"strconv"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

/* Parse status */
func ParseID(arg string) int {
	id, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Println("Invalid task ID:", arg)
		os.Exit(1)
	}
	return id
}
