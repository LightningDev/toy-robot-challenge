package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type LeftCommand struct {
	Name string
}

func NewLeftCommand(args []string) (robot.Command, error) {
	return LeftCommand{
		Name: "LEFT",
	}, nil
}

func (c LeftCommand) Execute(r *robot.Robot) error {
	r.Position.Rotate(-90) // Clockwise rotation

	return nil
}

func (c LeftCommand) GetName() string {
	return c.Name
}
