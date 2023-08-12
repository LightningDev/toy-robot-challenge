package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/command"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

func ParseCommand(robotCommand string, robot robot.Robot) (robot.Command, error) {
	// Regexp pattern for robot command: Command Args[]
	pattern := `^(?i)([A-Za-z]+)(?:\s+(\S.*))?`
	regex := regexp.MustCompile(pattern)

	// Extract command name and arguments
	matches := regex.FindStringSubmatch(robotCommand)
	if matches == nil {
		return nil, &errors.ValidationError{
			Command: "EMPTY",
			Err:     fmt.Errorf("Empty command"),
		}
	}

	commandName := strings.ToUpper(matches[1])
	args := strings.Split(matches[2], ",")

	// Skip if robot is not active
	if commandName != "PLACE" && !robot.Active {
		return nil, fmt.Errorf("robot is not active")
	}

	foundCommand, found := command.CommandList[commandName]
	if !found {
		return nil, &errors.ValidationError{
			Command: commandName,
			Err:     fmt.Errorf("%s is not a valid command", commandName),
		}
	}

	cmd, err := foundCommand(args)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
