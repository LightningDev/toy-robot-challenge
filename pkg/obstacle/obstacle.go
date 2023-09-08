package obstacle

type Obstacle struct {
	X int
	Y int
}

func New(x int, y int) *Obstacle {
	return &Obstacle{X: x, Y: y}
}
