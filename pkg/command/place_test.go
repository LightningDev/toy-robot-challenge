package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewPlaceCommand(t *testing.T) {
	tests := []struct {
		args       []string
		x          int
		y          int
		dir        position.Direction
		shouldFail bool
	}{
		{[]string{"1", "1", "NORTH"}, 1, 1, position.NORTH, false},
		{[]string{"1", "NORTH"}, 0, 0, position.NORTH, true},
		{[]string{}, 0, 0, position.EAST, true},
	}

	for _, test := range tests {
		cmd, err := NewPlaceCommand(test.args)
		if err != nil && !test.shouldFail {
			t.Fatalf("Expected no error but got: %v", err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("Expected error but got none")
		}

		if err == nil {
			placeCmd, ok := cmd.(PlaceCommand)
			if !ok {
				t.Fatalf("Expected PlaceCommand type, got: %T", cmd)
			}

			if placeCmd.Name != "PLACE" || placeCmd.X != test.x || placeCmd.Y != test.y || placeCmd.Facing != test.dir {
				t.Errorf("Expected command to match input but got different values")
			}
		}
	}
}

func TestExecutePlaceCommand(t *testing.T) {
	board := table.New(5, 5)
	tests := []struct {
		x          int
		y          int
		dir        position.Direction
		active     bool
		shouldFail bool
	}{
		{1, 1, position.EAST, true, false},
		{6, 6, position.WEST, false, true},
	}

	for _, test := range tests {
		r := &robot.Robot{}
		cmd := PlaceCommand{
			Name:   "PLACE",
			X:      test.x,
			Y:      test.y,
			Facing: test.dir,
		}

		err := cmd.Execute(r, *board)
		if err != nil && !test.shouldFail {
			t.Errorf("Expected no error but got: %v", err)
		}

		if err == nil && test.shouldFail {
			t.Fatalf("Expected error but got none")
		}

		if !test.shouldFail && (r.Position.X != test.x || r.Position.Y != test.y || r.Position.Direction != test.dir || r.Active != test.active) {
			t.Errorf("Expected robot properties to match input but got different values")
		}
	}
}

func TestGetNamePlaceCommand(t *testing.T) {
	cmd := PlaceCommand{
		Name: "REPORT",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
