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
	Long:  `Add a robot command`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
