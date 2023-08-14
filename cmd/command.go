/*
Copyright © 2023 Nhat Tran nhat1811@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/LightningDev/toy-robot-challenge/internal/generator"
	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Add a robot command",
	Long: `Generate a command file for the robot
Example:
  run add command JUMP

This will generate a file named jump.go & jump_test.go in the pkg/cmd directory.
`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		commandTemplate := &generator.CmdTemplate{
			AbsolutePath: wd,
		}

		for _, arg := range args {
			commandTemplate.CmdName = arg
			commandTemplate.Create()
		}
	},
}

func init() {
	addCmd.AddCommand(commandCmd)
}
