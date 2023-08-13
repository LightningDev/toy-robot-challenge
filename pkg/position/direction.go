package position

import "strings"

type Direction int

const (
	NORTH Direction = 0
	SOUTH Direction = 180
	EAST  Direction = 90
	WEST  Direction = 270
)

var directionList = map[Direction]string{
	NORTH: "NORTH",
	SOUTH: "SOUTH",
	EAST:  "EAST",
	WEST:  "WEST",
}

func (d Direction) String() string {
	if str, found := directionList[d]; found {
		return str
	}
	return "UNKNOWN"
}

func (d Direction) Degree() int {
	return int(d)
}

func StrToDirection(str string) Direction {
	for k, v := range directionList {
		if v == strings.ToUpper(str) {
			return k
		}
	}
	return -1
}

func DegreeToDirection(degree int) Direction {
	switch degree {
	case 0:
		return NORTH
	case 180:
		return SOUTH
	case 90:
		return EAST
	case 270:
		return WEST
	}

	return -1
}
