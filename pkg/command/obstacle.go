package command

import (
	"fmt"
	"strconv"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type ObstacleCommand struct {
	Name string
	X    int
	Y    int
}

func NewObstacleCommand(args []string) (robot.Command, error) {
	if len(args) != 2 {
		return nil, &errors.ValidationError{
			Command: "OBSTACLE",
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

	return ObstacleCommand{
		Name: "OBSTACLE",
		X:    newX,
		Y:    newY,
	}, nil
}

func (c ObstacleCommand) Execute(r *robot.Robot, t *table.Table) error {
	if (r.Position.X != c.X || r.Position.Y != c.Y) && t.IsValidPosition(c.X, c.Y) {
		t.Obstacles = append(t.Obstacles, *obstacle.New(c.X, c.Y))
	} else {
		return &errors.ValidationError{
			Command: c.Name,
			Err:     fmt.Errorf("invalid position: %+v, %+v", c.X, c.Y),
		}
	}

	return nil
}

func (c ObstacleCommand) GetName() string {
	return c.Name
}
