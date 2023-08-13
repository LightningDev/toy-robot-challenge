package command

import (
	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
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
	err := r.Position.Forward()
	if err != nil {
		return &errors.ValidationError{
			Command: c.Name,
			Err:     err,
		}
	}

	return nil
}

func (c MoveCommand) GetName() string {
	return c.Name
}
