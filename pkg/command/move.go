package command

import (
	"fmt"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type MoveCommand struct {
	Name string
}

func NewMoveCommand(args []string) (robot.Command, error) {
	return MoveCommand{
		Name: "MOVE",
	}, nil
}

func (c MoveCommand) Execute(r *robot.Robot) error {
	newX := r.CurrentX
	newY := r.CurrentY

	switch r.Facing {
	case position.NORTH:
		newY++
	case position.SOUTH:
		newY--
	case position.EAST:
		newX++
	case position.WEST:
		newX--
	}

	if !table.IsValidPosition(newX, newY) {
		return &errors.ValidationError{
			Command: c.Name,
			Err:     fmt.Errorf("invalid position"),
		}
	}
	fmt.Println("MOVE: ", newX, newY, r.Facing.String())
	r.CurrentX = newX
	r.CurrentY = newY

	return nil
}
