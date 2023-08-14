package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewLeftCommand(t *testing.T) {
	cmd, err := NewLeftCommand(nil)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	LeftCmd, ok := cmd.(LeftCommand)
	if !ok {
		t.Fatalf("Expected LeftCommand type, got: %T", cmd)
	}

	if LeftCmd.Name != "LEFT" {
		t.Errorf("Expected command name to be 'LEFT', got: %s", LeftCmd.Name)
	}
}

func TestExecuteLeftCommand(t *testing.T) {
	board := table.New(5, 5)
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	cmd := LeftCommand{
		Name: "LEFT",
	}

	err := cmd.Execute(r, *board)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if r.Position.Direction != position.WEST {
		t.Errorf("Expected robot direction to be 'WEST', got: %s", r.Position.Direction)
	}
}

func TestGetNameLeftCommand(t *testing.T) {
	cmd := LeftCommand{
		Name: "LEFT",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
