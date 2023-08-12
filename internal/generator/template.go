package generator

func CommandTemplate() []byte {
	return []byte(`package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type {{ .CmdName }}Command struct {
	Name string
}

func New{{ .CmdName }}Command(args []string) (robot.Command, error) {
	return {{ .CmdName }}Command{
		Name: "{{ toupper .CmdName }}",
	}, nil
}

func (c {{ .CmdName }}Command) Execute(r *robot.Robot) error {
	return nil
}
`)
}

func CommandListTemplate() []byte {
	return []byte(`package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

var CommandList = map[string]func([]string) (robot.Command, error){
	"PLACE":  NewPlaceCommand,
	"REPORT": NewReportCommand,
	"MOVE":   NewMoveCommand,
}
`)
}
