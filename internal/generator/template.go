package generator

func CommandTemplate() []byte {
	return []byte(`package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type {{ totitle .CmdName }}Command struct {
	Name string
}

func New{{ totitle .CmdName }}Command(args []string) (robot.Command, error) {
	return {{ totitle .CmdName }}Command{
		Name: "{{ toupper .CmdName }}",
	}, nil
}

func (c {{ totitle .CmdName }}Command) Execute(r *robot.Robot) error {
	return nil
}

func (c {{ totitle .CmdName }}Command) GetName() string {
	return c.Name
}
`)
}

func CommandTestTemplate() []byte {
	return []byte(`package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

func TestNew{{ totitle .CmdName }}Command(t *testing.T) {
	cmd, err := New{{ totitle .CmdName }}Command(nil)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	{{ tolower .CmdName }}Cmd, ok := cmd.({{ totitle .CmdName }}Command)
	if !ok {
		t.Fatalf("Expected {{ totitle .CmdName }}Command type, got: %T", cmd)
	}

	if {{ tolower .CmdName }}Cmd.Name != "{{ toupper .CmdName }}" {
		t.Errorf("Expected command name to be '{{ toupper .CmdName }}', got: %s", {{ tolower .CmdName }}Cmd.Name)
	}
}

func TestExecute{{ totitle .CmdName }}Command(t *testing.T) {
	r := &robot.Robot{}

	cmd := {{ totitle .CmdName }}Command{
		Name: "{{ toupper .CmdName }}",
	}

	err := cmd.Execute(r)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}
}

func TestGetName{{ totitle .CmdName }}Command(t *testing.T) {
	cmd := {{ totitle .CmdName }}Command{
		Name: "{{ toupper .CmdName }}",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
`)
}

func CommandListTemplate() []byte {
	return []byte(`package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

var CommandList = map[string]func([]string) (robot.Command, error){
	{{- range .Commands }}
	"{{ . }}": New{{ totitle . }}Command,
	{{- end }}
}
`)
}
