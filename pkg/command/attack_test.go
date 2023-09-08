package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewAttackCommand(t *testing.T) {
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
		cmd, err := NewAttackCommand(test.args)
		if err != nil && !test.shouldFail {
			t.Fatalf("Expected no error but got: %v", err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("Expected error but got none")
		}

		if err == nil {
			atkCmd, ok := cmd.(AttackCommand)
			if !ok {
				t.Fatalf("Expected AttackCommand type, got: %T", cmd)
			}

			if atkCmd.Name != "ATTACK" || atkCmd.X != test.x || atkCmd.Y != test.y {
				t.Errorf("Expected command to match input but got different values")
			}
		}
	}
}

func TestExecuteAttackCommand(t *testing.T) {
	board := table.New(10, 10, []obstacle.Obstacle{
		{X: 3, Y: 5},
		{X: 6, Y: 1},
		{X: 6, Y: 2},
		{X: 7, Y: 3},
		{X: 8, Y: 2},
		{X: 3, Y: 3},
		{X: 1, Y: 0},
	})
	r := &robot.Robot{
		Position: position.Position{X: 7, Y: 2, Direction: position.NORTH},
	}

	tests := []struct {
		x          int
		y          int
		shouldFail bool
	}{
		{0, 0, false},
		{1, 0, true},
	}

	for _, test := range tests {
		cmd := AttackCommand{
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

func TestGetNameAttackCommand(t *testing.T) {
	cmd := AttackCommand{
		Name: "ATTACK",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
