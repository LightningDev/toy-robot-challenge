package table

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
)

const (
	defaultBoardX = 5
	defaultBoardY = 5
)

type Table struct {
	Width     int
	Height    int
	Obstacles []obstacle.Obstacle
}

func New(width int, height int, obstacles []obstacle.Obstacle) *Table {
	if width <= 0 {
		width = defaultBoardX
	}
	if height <= 0 {
		height = defaultBoardY
	}

	return &Table{Width: width, Height: height, Obstacles: obstacles}
}

func (t *Table) IsValidPosition(x, y int) bool {
	for _, obstacle := range t.Obstacles {
		if x == obstacle.X && y == obstacle.Y {
			return false
		}
	}
	return x >= 0 && x < t.Width && y >= 0 && y < t.Height
}
