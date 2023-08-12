package position

import (
	"errors"
	"fmt"

	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type Position struct {
	X         int
	Y         int
	Direction Direction
}

func (p *Position) Rotate(degree int) {
	if degree < 0 {
		degree += 360
	}
	newDirection := (p.Direction.Degree() + degree) % 360

	p.Direction = DegreeToDirection(newDirection)
}

func (p *Position) Forward() error {
	return move(p, 1)
}

func (p *Position) Backward() error {
	return move(p, -1) // Backward is just a negative forward
}

func move(p *Position, step int) error {
	newX := p.X
	newY := p.Y

	switch p.Direction {
	case NORTH:
		newY += step
	case SOUTH:
		newY -= step
	case EAST:
		newX += step
	case WEST:
		newX -= step
	}

	if !table.IsValidPosition(newX, newY) {
		return errors.New("invalid position")
	}

	p.X = newX
	p.Y = newY

	return nil
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d,%s", p.X, p.Y, p.Direction)
}
