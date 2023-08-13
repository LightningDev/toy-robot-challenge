package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
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
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	cmd := RightCommand{
		Name: "RIGHT",
	}

	err := cmd.Execute(r)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if r.Position.Direction != position.EAST {
		t.Errorf("Expected robot direction to be 'EAST', got: %s", r.Position.Direction)
	}
}
