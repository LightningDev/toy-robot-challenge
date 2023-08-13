package command

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
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
	r := &robot.Robot{
		Position: position.Position{X: 2, Y: 2, Direction: position.NORTH},
	}

	cmd := ReportCommand{
		Name: "REPORT",
	}

	err := cmd.Execute(r)
	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}
}