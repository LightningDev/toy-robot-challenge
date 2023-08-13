package command

import (
	"fmt"
	"strconv"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type PlaceCommand struct {
	Name   string
	X      int
	Y      int
	Facing position.Direction
}

func NewPlaceCommand(args []string) (robot.Command, error) {
	if len(args) != 3 {
		return nil, &errors.ValidationError{
			Command: "PLACE",
			Err:     fmt.Errorf("invalid arguments"),
		}
	}

	newX, _ := strconv.Atoi(args[0])
	newY, _ := strconv.Atoi(args[1])
	facing := position.StrToDirection(args[2])

	return PlaceCommand{
		Name:   "PLACE",
		X:      newX,
		Y:      newY,
		Facing: facing,
	}, nil
}

func (c PlaceCommand) Execute(r *robot.Robot) error {
	if !table.IsValidPosition(c.X, c.Y) {
		r.Active = false || r.Active
		return &errors.ValidationError{
			Command: c.Name,
			Err:     fmt.Errorf("invalid position"),
		}
	}

	r.Position.X = c.X
	r.Position.Y = c.Y
	r.Position.Direction = c.Facing
	r.Active = true
	return nil
}
