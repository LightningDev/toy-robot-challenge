/*
Copyright © 2023 Nhat Tran nhat1811@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new logic too app",
	Long: `This command will add a new logic to the app

Current supported logic:
- command: add a new command for the robot
`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
