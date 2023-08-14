/*
Copyright © 2023 Nhat Tran nhat1811@gmail.com
*/
package cmd

import (
	"os"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "The application is a simulation of a toy robot moving on a square tabletop.",
	Long: `Start playing the toy robot game by running:
  run play
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&errors.Debug, "debug", "d", false, "Toggle debug mode to debug the application")
}
