/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/internal/parser"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/spf13/cobra"
)

// File name
var filename string

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start playing toy robot game",
	Long: `You have a toy robot on a table top, a grid of 5 x 5 units, there are no obstructions.
You can issue commands to your robot allowing it to roam around the table top. But be careful, don't let it fall off!

Available commands:
- PLACE X,Y,FACING: X and Y is coordiantes of the robot, FACING is direction of the robot (NORTH, SOUTH, EAST, WEST)
- MOVE: move the toy robot one unit forward in the direction it is currently facing
- LEFT: rotate the robot 90 degrees in the specified direction without changing the position of the robot
- RIGHT: rotate the robot 90 degrees in the specified direction without changing the position of the robot
- REPORT: announce the X, Y, and orientation of the robot on the table
`,
	Run: inputCommand,
}

// Allow user to input command from the console
func inputCommand(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Long)
	fmt.Println("Each command must be on a separate line.")
	fmt.Println("Please start entering commands. Enter EXIT to exit.")

	// Set up a signal handler to capture Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	// Input loop
	robot := &robot.Robot{}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = input[:len(input)-1] // Remove the newline character
		if input == "EXIT" {
			break
		}

		command, err := parser.ParseCommand(input, *robot)

		if err != nil {
			//fmt.Println(err.Error())
			continue
		}

		err = robot.Do(command)
		if err != nil {
			errors.HandleError(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	playCmd.PersistentFlags().StringVarP(&filename, "file", "f", "", "Read input from a file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
