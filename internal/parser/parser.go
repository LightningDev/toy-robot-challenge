package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/pkg/command"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

var regex = regexp.MustCompile(`^(?i)([A-Za-z]+)(?:\s+(\S.*))?`)

func ParseCommand(robotCommand string, robot robot.Robot) (robot.Command, error) {
	// Extract command name and arguments
	matches := regex.FindStringSubmatch(robotCommand)
	if matches == nil {
		return nil, &errors.ValidationError{
			Command: "EMPTY",
			Err:     fmt.Errorf("empty command"),
		}
	}

	commandName := strings.ToUpper(matches[1])
	args := strings.Split(strings.TrimSpace(matches[2]), ",")

	// Skip if robot is not active
	if commandName != "PLACE" && !robot.Active {
		return nil, fmt.Errorf("please place the robot on the board first")
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
