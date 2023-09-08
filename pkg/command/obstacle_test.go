package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewObstacleCommand(t *testing.T) {
	tests := []struct {
		args       []string
		x          int
		y          int
		shouldFail bool
	}{
		{[]string{"1", "1"}, 1, 1, false},
		{[]string{}, 0, 0, true},
	}

	for _, test := range tests {
		cmd, err := NewObstacleCommand(test.args)
		if err != nil && !test.shouldFail {
			t.Fatalf("Expected no error but got: %v", err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("Expected error but got none")
		}

		if err == nil {
			obsCmd, ok := cmd.(ObstacleCommand)
			if !ok {
				t.Fatalf("Expected ObstacleCommand type, got: %T", cmd)
			}

			if obsCmd.Name != "OBSTACLE" || obsCmd.X != test.x || obsCmd.Y != test.y {
				t.Errorf("Expected command to match input but got different values")
			}
		}
	}
}

func TestExecuteObstacleCommand(t *testing.T) {
	board := table.New(5, 5, []obstacle.Obstacle{})
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	tests := []struct {
		x          int
		y          int
		shouldFail bool
	}{
		{1, 1, false},
		{2, 2, true},
	}

	for _, test := range tests {
		cmd := ObstacleCommand{
			Name: "OBSTACLE",
			X:    test.x,
			Y:    test.y,
		}

		err := cmd.Execute(r, board)
		if err != nil && !test.shouldFail {
			t.Errorf("Expected no error but got: %v", err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("Expected error but got none")
		}
	}
}

func TestGetNameObstacleCommand(t *testing.T) {
	cmd := ObstacleCommand{
		Name: "OBSTACLE",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
