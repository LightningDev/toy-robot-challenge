package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

var CommandList = map[string]func([]string) (robot.Command, error){
	"PLACE": NewPlaceCommand,
	"REPORT": NewReportCommand,
	"MOVE": NewMoveCommand,
	"LEFT": NewLeftCommand,
	"RIGHT": NewRightCommand,
	"OBSTACLE": NewObstacleCommand,
	"ATTACK": NewAttackCommand,
}
