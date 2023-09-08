package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewRightCommand(t *testing.T) {
	cmd, err := NewRightCommand(nil)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	rightCmd, ok := cmd.(RightCommand)
	if !ok {
		t.Fatalf("Expected RightCommand type, got: %T", cmd)
	}

	if rightCmd.Name != "RIGHT" {
		t.Errorf("Expected command name to be 'RIGHT', got: %s", rightCmd.Name)
	}
}

func TestExecuteRightCommand(t *testing.T) {
	board := table.New(5, 5, []obstacle.Obstacle{})
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	cmd := RightCommand{
		Name: "RIGHT",
	}

	err := cmd.Execute(r, board)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if r.Position.Direction != position.EAST {
		t.Errorf("Expected robot direction to be 'EAST', got: %s", r.Position.Direction)
	}
}

func TestGetNameRightCommand(t *testing.T) {
	cmd := RightCommand{
		Name: "RIGHT",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
