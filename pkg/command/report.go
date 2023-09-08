package command

import (
	"fmt"

	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type ReportCommand struct {
	Name string
}

func NewReportCommand(args []string) (robot.Command, error) {
	return ReportCommand{
		Name: "REPORT",
	}, nil
}

func (c ReportCommand) Execute(r *robot.Robot, t *table.Table) error {
	fmt.Printf("Output: %s\n", r.Position)
	return nil
}

func (c ReportCommand) GetName() string {
	return c.Name
}
