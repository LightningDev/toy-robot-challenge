package command

import (
	"fmt"
	"strconv"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/path"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type AttackCommand struct {
	Name string
	X    int
	Y    int
}

func NewAttackCommand(args []string) (robot.Command, error) {
	if len(args) != 2 {
		return nil, &errors.ValidationError{
			Command: "ATTACK",
			Err:     fmt.Errorf("invalid arguments"),
		}
	}

	newX, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, &errors.ValidationError{
			Command: "PLACE",
			Err:     fmt.Errorf("invalid X position"),
		}
	}
	newY, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, &errors.ValidationError{
			Command: "PLACE",
			Err:     fmt.Errorf("invalid Y position"),
		}
	}

	return AttackCommand{
		Name: "ATTACK",
		X:    newX,
		Y:    newY,
	}, nil
}

func (c AttackCommand) Execute(r *robot.Robot, t *table.Table) error {
	if !t.IsValidPosition(c.X, c.Y) {
		return &errors.ValidationError{
			Command: "ATTACK",
			Err:     fmt.Errorf("invalid position to attack: %+v, %+v", c.X, c.Y),
		}
	}
	foundPath := path.Find(r.Position, position.Position{X: c.X, Y: c.Y}, *t)
	if len(foundPath) == 0 {
		return &errors.ValidationError{
			Command: "ATTACK",
			Err:     fmt.Errorf("no path found to attack: %+v, %+v", c.X, c.Y),
		}
	}

	fmt.Printf("Possible Path To Attack at: %+v, %+v\n", c.X, c.Y)
	for index, node := range foundPath {
		fmt.Printf("Step %+v: %s\n", index, node)
	}
	return nil
}

func (c AttackCommand) GetName() string {
	return c.Name
}
