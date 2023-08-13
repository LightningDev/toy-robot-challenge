package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CmdTemplate struct {
	CmdName      string
	AbsolutePath string
}

type CmdListTemplate struct {
	Commands []string
}

func (t *CmdTemplate) Create() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(t.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(t.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	// create pkg/command/<command_name>.go
	cmdFilePath := fmt.Sprintf("%s/pkg/command/%s.go", t.AbsolutePath, strings.ToLower(t.CmdName))
	err := createFile(t, cmdFilePath, CommandTemplate())
	if err != nil {
		return err
	}

	// create pkg/command/<command_name>_test.go
	cmdTestFilePath := fmt.Sprintf("%s/pkg/command/%s_test.go", t.AbsolutePath, strings.ToLower(t.CmdName))
	err = createFile(t, cmdTestFilePath, CommandTestTemplate())
	if err != nil {
		return err
	}

	// create pkg/command/command.go
	cmdListTemplate, err := updateCommandsJSON(t.CmdName, fmt.Sprintf("%s/config/command.json", t.AbsolutePath))
	if err != nil {
		return err
	}

	cmdListFilePath := fmt.Sprintf("%s/pkg/command/command.go", t.AbsolutePath)
	err = createFile(cmdListTemplate, cmdListFilePath, CommandListTemplate())
	if err != nil {
		return err
	}

	return nil
}

// Create a new file from template
func createFile(t any, path string, data []byte) error {
	cmdFile, err := os.Create(path)
	if err != nil {
		return err
	}

	defer cmdFile.Close()

	template := createTemplate("cmd", data)
	err = template.Execute(cmdFile, t)
	if err != nil {
		return err
	}

	return nil
}

// Create a template object from template byte data
func createTemplate(name string, data []byte) *template.Template {
	return template.Must(template.New(name).Funcs(template.FuncMap{
		"toupper": strings.ToUpper,
		"tolower": strings.ToLower,
		"totitle": cases.Title(language.English).String,
	}).Parse(string(data)))
}

// Update the command list config after adding a new one
func updateCommandsJSON(newCommand, pathToJSON string) (*CmdListTemplate, error) {
	newCommand = strings.ToUpper(newCommand)
	var cmdList CmdListTemplate

	data, err := os.ReadFile(pathToJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cmdList)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(cmdList.Commands, newCommand) {
		cmdList.Commands = append(cmdList.Commands, newCommand)

		updatedData, err := json.MarshalIndent(cmdList, "", "  ")
		if err != nil {
			return nil, err
		}

		return &cmdList, os.WriteFile(pathToJSON, updatedData, 0644)
	}

	return &cmdList, nil
}
