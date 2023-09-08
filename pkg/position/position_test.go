package position

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

var board = table.New(5, 5, []obstacle.Obstacle{})

func TestRotate(t *testing.T) {
	tests := []struct {
		startDirection Direction
		degree         int
		expected       Direction
	}{
		{NORTH, 90, EAST},
		{NORTH, 180, SOUTH},
		{EAST, -90, NORTH},
		{EAST, 270, NORTH},
	}

	for _, tt := range tests {
		pos := &Position{Direction: tt.startDirection}
		pos.Rotate(tt.degree)
		if pos.Direction != tt.expected {
			t.Errorf("Rotation of %v by %d degrees, expected %v but got %v", tt.startDirection, tt.degree, tt.expected, pos.Direction)
		}
	}
}

func TestForward(t *testing.T) {
	tests := []struct {
		startX, startY int
		direction      Direction
		expectedError  bool
		expectedX      int
		expectedY      int
	}{
		{2, 2, NORTH, false, 2, 3},
		{4, 4, NORTH, true, 4, 4},
		{2, 2, SOUTH, false, 2, 1},
	}

	for _, tt := range tests {
		pos := &Position{X: tt.startX, Y: tt.startY, Direction: tt.direction}
		err := pos.Forward(*board, 1)
		if (err != nil) != tt.expectedError {
			t.Errorf("For position (%d,%d) facing %v, expected error to be %v but got %v", tt.startX, tt.startY, tt.direction, tt.expectedError, err)
		}
		if pos.X != tt.expectedX || pos.Y != tt.expectedY {
			t.Errorf("For position (%d,%d) facing %v, expected position after moving forward to be (%d,%d) but got (%d,%d)", tt.startX, tt.startY, tt.direction, tt.expectedX, tt.expectedY, pos.X, pos.Y)
		}
	}
}

func TestBackward(t *testing.T) {
	tests := []struct {
		startX, startY int
		direction      Direction
		expectedError  bool
		expectedX      int
		expectedY      int
	}{
		{2, 2, NORTH, false, 2, 1},
		{4, 4, WEST, true, 4, 4},
		{0, 4, SOUTH, true, 0, 4},
	}

	for _, tt := range tests {
		pos := &Position{X: tt.startX, Y: tt.startY, Direction: tt.direction}
		err := pos.Backward(*board, 1)
		if (err != nil) != tt.expectedError {
			t.Errorf("For position (%d,%d) facing %v, expected error to be %v but got %v", tt.startX, tt.startY, tt.direction, tt.expectedError, err)
		}
		if pos.X != tt.expectedX || pos.Y != tt.expectedY {
			t.Errorf("For position (%d,%d) facing %v, expected position after moving forward to be (%d,%d) but got (%d,%d)", tt.startX, tt.startY, tt.direction, tt.expectedX, tt.expectedY, pos.X, pos.Y)
		}
	}
}

func TestPositionString(t *testing.T) {
	tests := []struct {
		x, y           int
		direction      Direction
		expectedString string
	}{
		{2, 2, NORTH, "2,2,NORTH"},
	}

	for _, tt := range tests {
		pos := Position{X: tt.x, Y: tt.y, Direction: tt.direction}
		str := pos.String()
		if str != tt.expectedString {
			t.Errorf("Expected position string to be %s but got %s", tt.expectedString, str)
		}
	}
}
