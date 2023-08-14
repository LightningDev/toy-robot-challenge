package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	testdata "github.com/LightningDev/toy-robot-challenge/test"
)

func TestPlayCommand(t *testing.T) {
	testCases, err := testdata.ParseJSON("../test/commands.json")
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	for _, testCase := range testCases {
		// Create a temporary file
		tempFile, err := os.CreateTemp("", "sample_command")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}

		// Write commands to the temporary file
		for _, command := range testCase.Commands {
			_, err := tempFile.WriteString(command + "\n")
			if err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}
		}

		// Close the file to ensure it's flushed
		tempFile.Close()

		// Backup the original stdout
		oldStdout := os.Stdout

		// Create a pipe. Writer side will be used to capture stdout
		// and Reader side to read the captured data
		reader, writer, _ := os.Pipe()
		os.Stdout = writer

		rootCmd.SetArgs([]string{"play", "-f", tempFile.Name()})
		rootCmd.Execute()

		// Close the writer side so that we can start reading from the reader side
		writer.Close()

		// Restore the original stdout
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(reader)
		actual := buf.String()

		for _, expected := range testCase.Output {
			if !strings.Contains(actual, expected) {
				t.Errorf("For commands %v, expected output %s not found in response. Got: %s", testCase.Commands, expected, actual)
			}
		}
	}
}
