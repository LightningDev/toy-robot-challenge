package command

import (
	"fmt"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
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
	newDegree := (r.Facing.Degree() - 90 + 360) % 360
	r.Facing = position.DegreeToDirection(newDegree)
	fmt.Println("LEFT: ", r.CurrentX, r.CurrentY, r.Facing.String())
	return nil
}
