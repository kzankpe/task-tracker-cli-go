package helper

import (
	"fmt"
	"os"
)

type task struct {
}

func CreateTaskFile(filePath string) error {
	fmt.Println("Creating file on the system")
	_, err := os.Create(filePath)

	if err != nil {
		return err
	}
	return nil
}

func TestTaskFile(filePath string) (bool, error) {
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
