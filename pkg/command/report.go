package command

import (
	"fmt"

	"github.com/LightningDev/toy-robot-challenge/internal/utils"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

type ReportCommand struct {
	Name string
}

func NewReportCommand(args []string) (robot.Command, error) {
	return ReportCommand{}, nil
}

func (c ReportCommand) Execute(r *robot.Robot) error {
	fmt.Printf("Output: %d,%d,%s", r.CurrentX, r.CurrentY, utils.DirectionToStr(r.Facing))
	return nil
}
