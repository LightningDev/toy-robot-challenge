package robot

import "github.com/LightningDev/toy-robot-challenge/pkg/position"

type Command interface {
	GetName() string
	Execute(r *Robot) error
}

type Robot struct {
	Position position.Position
	Active   bool
}

func (r *Robot) Do(command Command) error {
	return command.Execute(r)
}
