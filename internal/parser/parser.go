package parser

import (
	"regexp"
	"strings"

	"github.com/LightningDev/toy-robot-challenge/pkg/command"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

var commandList = map[string]func([]string) (robot.Command, error){
	"PLACE":  command.NewPlaceCommand,
	"REPORT": command.NewReportCommand,
}

func ParseCommand(robotCommand string) (robot.Command, error) {
	// Regexp pattern for robot command: Command Args[]
	pattern := `^([A-Z]+)(?:\s+(\S.*))?`
	regex := regexp.MustCompile(pattern)

	// Extract command name and arguments
	matches := regex.FindStringSubmatch(robotCommand)
	commandName := matches[1]
	args := strings.Split(matches[2], ",")
	foundCommand, err := commandList[commandName](args)

	if err != nil {
		return nil, err
	}

	return foundCommand, nil
}
