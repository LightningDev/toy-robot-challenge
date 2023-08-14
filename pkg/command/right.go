package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type RightCommand struct {
	Name string
}

func NewRightCommand(args []string) (robot.Command, error) {
	return RightCommand{
		Name: "RIGHT",
	}, nil
}

func (c RightCommand) Execute(r *robot.Robot, t table.Table) error {
	r.Position.Rotate(90) // Counter clockwise rotation

	return nil
}

func (c RightCommand) GetName() string {
	return c.Name
}
