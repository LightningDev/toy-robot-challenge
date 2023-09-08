package robot

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type Command interface {
	GetName() string
	Execute(r *Robot, t *table.Table) error
}

type Robot struct {
	Position position.Position
	Active   bool
}

func (r *Robot) Do(command Command, t *table.Table) error {
	return command.Execute(r, t)
}
