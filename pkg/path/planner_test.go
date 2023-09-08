package path

import (
	"reflect"
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestFind(t *testing.T) {
	// Test data
	tableWithObstacles := table.New(5, 5, []obstacle.Obstacle{
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
	})
	startPos := position.Position{X: 0, Y: 0}
	endPos := position.Position{X: 4, Y: 4}

	result := Find(startPos, endPos, *tableWithObstacles)

	expectedPath := []position.Position{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 1},
		{X: 4, Y: 2},
		{X: 4, Y: 3},
		{X: 4, Y: 4},
	}

	// Check the result
	if !reflect.DeepEqual(result, expectedPath) {
		t.Errorf("Expected %v but got %v", expectedPath, result)
	}
}
