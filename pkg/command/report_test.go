package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

func TestNewReportCommand(t *testing.T) {
	cmd, err := NewReportCommand(nil)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	reportCmd, ok := cmd.(ReportCommand)
	if !ok {
		t.Fatalf("Expected ReportCommand type, got: %T", cmd)
	}

	if reportCmd.Name != "REPORT" {
		t.Errorf("Expected command name to be 'REPORT', got: %s", reportCmd.Name)
	}
}

func TestExecuteReportCommand(t *testing.T) {
	board := table.New(5, 5)
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	cmd := ReportCommand{
		Name: "REPORT",
	}

	err := cmd.Execute(r, *board)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}
}

func TestGetNameReportCommand(t *testing.T) {
	cmd := ReportCommand{
		Name: "REPORT",
	}

	name := cmd.GetName()
	if name != cmd.Name {
		t.Fatalf("Expected command name is %s but got %s", cmd.Name, name)
	}
}
