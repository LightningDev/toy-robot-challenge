package mockup

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CommandTemplate() []byte {
	return []byte(`
package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type {{ .CmdName }}Command struct {
	Name string
}

func New{{ .CmdName }}Command(args []string) (robot.Command, error) {
	return PlaceCommand{
		Name: "{{ .CmdName }}",
	}, nil
}

func (c {{ .CmdName }}Command) Execute(r *robot.Robot) error {
	return nil
}
`)
}

type Template struct {
	CmdName      string
	AbsolutePath string
}

func (t *Template) Create() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(t.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(t.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	// create pkg/command/<command_name>.go
	commandFile, err := os.Create(fmt.Sprintf("%s/pkg/command/%s.go", t.AbsolutePath, strings.ToLower(t.CmdName)))
	if err != nil {
		return err
	}
	defer commandFile.Close()

	commandTemplate := template.Must(template.New(t.CmdName).Parse(string(CommandTemplate())))

	t.CmdName = cases.Title(language.English).String(t.CmdName)
	err = commandTemplate.Execute(commandFile, t)
	if err != nil {
		return err
	}

	return nil
}
