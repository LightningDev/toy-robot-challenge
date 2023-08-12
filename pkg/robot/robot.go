package robot

import "github.com/LightningDev/toy-robot-challenge/pkg/position"

type Command interface {
	Execute(r *Robot) error
}

type Robot struct {
	CurrentX int
	CurrentY int
	Facing   position.Direction
}

func (r *Robot) Do(command Command) error {
	return command.Execute(r)
}
