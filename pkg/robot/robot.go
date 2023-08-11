package robot

type Direction int
type Rotation string

type Command interface {
	Execute(r *Robot) error
}

const (
	NORTH Direction = 0
	SOUTH Direction = 90
	EAST  Direction = 180
	WEST  Direction = 270
)

const (
	LEFT  Rotation = "LEFT"
	RIGHT Rotation = "RIGHT"
)

type Robot struct {
	CurrentX int
	CurrentY int
	Facing   Direction
}

func (r *Robot) Do(command Command) error {
	return command.Execute(r)
}
