package testdata

import (
	"encoding/json"
	"os"
)

type TestCommand struct {
	Commands []string `json:"commands"`
	Output   []string `json:"output"`
}

func ParseJSON(filename string) ([]TestCommand, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var testCases []TestCommand
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&testCases); err != nil {
		return nil, err
	}

	return testCases, nil
}
