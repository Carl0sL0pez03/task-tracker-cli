package task

import (
	"encoding/json"
	"os"
)

/* Load Task */
func LoadTask(filename string) (Tasks, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return Tasks{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks Tasks
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return tasks, nil
}

/* Save Task */
func SaveTask(filename string, tasks Tasks) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}
