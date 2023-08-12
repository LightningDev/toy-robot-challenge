package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type RightCommand struct {
	Name string
}

func NewRightCommand(args []string) (robot.Command, error) {
	return RightCommand{
		Name: "RIGHT",
	}, nil
}

func (c RightCommand) Execute(r *robot.Robot) error {
	newDegree := (r.Facing.Degree() + 90) % 360
	r.Facing = position.DegreeToDirection(newDegree)

	return nil
}
