package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func fileLocation() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	taskfilename := "task-" + user.Username + ".json"
	filePath := dirname + "/" + taskfilename

	return filePath
}

func CreateTaskFile() error {
	//Getting all variables from the system
	filepath := fileLocation()
	taskList := []Task{}
	jtaskList, _ := json.MarshalIndent(taskList, "", " ")
	msg := fmt.Sprintf("Creating file on the system at %s", filepath)
	fmt.Println(msg)
	f, err := os.Create(filepath)

	if err != nil {
		return err
	}
	defer f.Close() // Close the file
	fmt.Println("Initializing the document")
	err = os.WriteFile(filepath, jtaskList, 0644)
	if err != nil {
		fmt.Println("Error writing Json data to the file")
		return err
	}
	return nil
}

func TestTaskFile() (bool, error) {

	filePath := fileLocation()

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%q is a directory, not a file", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	return true, nil

}

func GetTaskFileContent() ([]Task, error) {
	filepath := fileLocation()
	var taskList []Task
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &taskList)

	if err != nil {
		return nil, err
	}
	return taskList, nil
}

func SaveTaskfile(tasklist []byte) error {
	filepath := fileLocation()
	err := os.WriteFile(filepath, tasklist, 0644)
	if err != nil {
		return err
	}
	return nil
}


func RemoveTask(tasks []Task, id int) []Task{
	for i, task := range tasks {
        if task.Id == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            break
        }
    }
    return tasks
}