package command

import (
	"fmt"

	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type ReportCommand struct {
	Name string
}

func NewReportCommand(args []string) (robot.Command, error) {
	return ReportCommand{
		Name: "REPORT",
	}, nil
}

func (c ReportCommand) Execute(r *robot.Robot) error {
	fmt.Printf("Output: %d,%d,%s\n", r.CurrentX, r.CurrentY, r.Facing.String())
	return nil
}
