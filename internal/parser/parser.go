package parser

import (
	"regexp"
	"strings"

	"github.com/LightningDev/toy-robot-challenge/pkg/command"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

func ParseCommand(robotCommand string) (robot.Command, error) {
	// Regexp pattern for robot command: Command Args[]
	pattern := `^(?i)([A-Za-z]+)(?:\s+(\S.*))?`
	regex := regexp.MustCompile(pattern)

	// Extract command name and arguments
	matches := regex.FindStringSubmatch(robotCommand)
	commandName := strings.ToUpper(matches[1])
	args := strings.Split(matches[2], ",")
	foundCommand, err := command.CommandList[commandName](args)

	if err != nil {
		return nil, err
	}

	return foundCommand, nil
}
