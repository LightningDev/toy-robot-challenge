package utils

import (
	"errors"

	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

var directionList = map[robot.Direction]string{
	robot.NORTH: "NORTH",
	robot.SOUTH: "SOUTH",
	robot.EAST:  "EAST",
	robot.WEST:  "WEST",
}

func StrToDirection(str string) (robot.Direction, error) {
	for dir, value := range directionList {
		if value == str {
			return dir, nil
		}
	}
	return -1, errors.New("invalid direction string")
}

func DirectionToStr(direction robot.Direction) string {
	if str, found := directionList[direction]; found {
		return str
	}
	return "UNKNOWN"
}
