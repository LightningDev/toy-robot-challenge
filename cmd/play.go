/*
Copyright © 2023 Nhat Tran nhat1811@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/LightningDev/toy-robot-challenge/internal/errors"
	"github.com/LightningDev/toy-robot-challenge/internal/parser"
	"github.com/LightningDev/toy-robot-challenge/pkg/obstacle"
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
	"github.com/spf13/cobra"
)

// File name
var filename string

// Table size and number of obstacles
var width, height int

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start playing toy robot game",
	Long: `You have a toy robot on a table top, a grid of nxn units, there are no obstructions.
You can issue commands to your robot allowing it to roam around the table top. But be careful, don't let it fall off!

Available commands:
- PLACE X,Y,FACING: X and Y is coordiantes of the robot, FACING is direction of the robot (NORTH, SOUTH, EAST, WEST)
- MOVE: move the toy robot one unit forward in the direction it is currently facing
- LEFT: rotate the robot 90 degrees in the specified direction without changing the position of the robot
- RIGHT: rotate the robot 90 degrees in the specified direction without changing the position of the robot
- REPORT: announce the X, Y, and orientation of the robot on the table
- OBSTACLE X,Y: place the obstacle on the table at position x and y
`,
	Run: inputCommand,
}

// Allow user to input command from the console
func inputCommand(cmd *cobra.Command, args []string) {
	robot := &robot.Robot{}
	board := table.New(width, height, []obstacle.Obstacle{})

	var reader *bufio.Reader

	if filename != "" {
		// If filename is provided, read commands from file
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", filename, err)
			return
		}
		defer file.Close()

		reader = bufio.NewReader(file)
	} else {
		// Else read from standard input
		boardSize := fmt.Sprintf("%d x %d", width, height)
		fmt.Println(strings.Replace(cmd.Long, "nxn", boardSize, 1))
		fmt.Println("Each command must be on a separate line.")
		fmt.Println("Please start entering commands. <Ctrl + c> to exit.")

		// Set up a signal handler to capture Ctrl+C
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigChan
			fmt.Println("\nExiting...")
			os.Exit(0)
		}()

		reader = bufio.NewReader(os.Stdin)
	}

	// Input loop
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSuffix(input, "\n")
		if err == io.EOF && input == "" {
			break
		}

		command, err := parser.ParseCommand(input, *robot)
		if err != nil {
			errors.HandleError(err)
			continue
		}

		err = robot.Do(command, board)
		if err != nil {
			errors.HandleError(err)
			continue
		}
	}
}

func init() {
	rootCmd.AddCommand(playCmd)

	playCmd.Flags().StringVarP(&filename, "file", "f", "", "Read input from a file")
	playCmd.Flags().IntVar(&width, "width", 5, "Width of the table")
	playCmd.Flags().IntVar(&height, "height", 5, "Height of the table")
}
