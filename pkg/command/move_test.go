package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewMoveCommand(t *testing.T) {
	cmd, err := NewMoveCommand(nil)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	MoveCmd, ok := cmd.(MoveCommand)
	if !ok {
		t.Fatalf("Expected MoveCommand type, got: %T", cmd)
	}

	if MoveCmd.Name != "MOVE" {
		t.Errorf("Expected command name to be 'MOVE', got: %s", MoveCmd.Name)
	}
}

func TestExecuteMoveCommand(t *testing.T) {
	board := table.New(5, 5, []obstacle.Obstacle{})
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	cmd := MoveCommand{
		Name: "MOVE",
	}

	err := cmd.Execute(r, board)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if r.Position.Y != 3 {
		t.Errorf("Expected robot Y position to be 3, got: %d", r.Position.Y)
	}

	if r.Position.X != 2 {
		t.Errorf("Expected robot X position to be 2, got: %d", r.Position.X)
	}

	if r.Position.Direction != position.NORTH {
		t.Errorf("Expected robot direction to be 'NORTH', got: %s", r.Position.Direction)
	}
}

func TestGetNameMoveCommand(t *testing.T) {
	cmd := MoveCommand{
		Name: "MOVE",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
