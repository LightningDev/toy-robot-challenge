package command

import (
	"github.com/LightningDev/toy-robot-challenge/internal/errors"
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

func (c MoveCommand) Execute(r *robot.Robot, t table.Table) error {
	err := r.Position.Forward(t)
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
