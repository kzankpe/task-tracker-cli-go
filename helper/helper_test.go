package helper

import (
	"os"
	"testing"
)

func TestCreateTaskFile(t *testing.T) {
	err := CreateTaskFile()
	if err != nil {
		t.Fatalf("Expected no error, fot %v", err)
	}

	filepath := fileLocation()
	defer os.Remove(filepath)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Fatalf("Expected file to be created at %s, but it does not exist", filepath)
	}

}

func TestContains(t *testing.T) {
	tests := []struct {
		table  [3]string
		elemet string
		want   bool
	}{
		{[3]string{"bread", "fish", "chicken"}, "fruit", false},
		{[3]string{"bread", "fish", "chicken"}, "fish", true},
		{[3]string{"fruit", "fish", "chicken"}, "fruit", true},
		{[3]string{"bread", "fish", "chicken"}, "", false},
	}

	for _, el := range tests {
		t.Run(el.elemet, func(t *testing.T) {
			res := Contains(el.table, el.elemet)
			if res != el.want {
				t.Errorf("Contains(%v, %q) = %v; want %v", el.table, el.elemet, res, el.want)
			}
		})
	}
}
