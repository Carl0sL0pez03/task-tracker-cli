package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks []Task

/* List Task */
func ListTasks(tasks Tasks, status string) {
	for _, task := range tasks {

		if status == "" || task.Status == status {
			fmt.Printf("ID: %d | Description: %s | Status: %s | CreatedAt: %s | UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
		}

	}
}

/* Add Task */
func AddTask(tasks *Tasks, description string) {
	newID := len(*tasks) + 1
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	*tasks = append(*tasks, newTask)
	fmt.Printf("Task added successfully (ID: %d)\n", newID)
}

/* Mark Task */
func MarkTask(tasks *Tasks, id int, status string) {
	for i := range *tasks {

		if (*tasks)[i].ID == id {
			(*tasks)[i].Status = status
			(*tasks)[i].UpdatedAt = time.Now()
			fmt.Printf("Task %d marked as %s\n", id)
			return
		}

	}
	fmt.Printf("Task with ID %d not found\n", id)
}

/* Update Task */
func UpdateTask(tasks *Tasks, id int, description string) {
	for i := range *tasks {

		if (*tasks)[i].ID == id {
			(*tasks)[i].Description = description
			(*tasks)[i].UpdatedAt = time.Now()
			fmt.Printf("Task %d updated successfully\n", id)
			return
		}

	}
	fmt.Printf("Task with ID %d not found\n", id)
}

/* Delete Task */
func DeleteTask(tasks *Tasks, id int) {
	for i := range *tasks {

		if (*tasks)[i].ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Printf("Task %d deleted successfully\n", id)
			return
		}

	}
	fmt.Printf("Task with ID %d not found\n", id)
}
