# Toy Robot Challenge
This project is an implementation of the Toy Robot simulation, allowing a toy robot to roam around a 5x5 or nxn (where n is any integer number) grid and execute a series of commands.

## Introduction
- The application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units.
- There are no other obstructions on the table surface.
- The robot is free to roam around the surface of the table, but must be prevented from falling to destruction. Any movement
  that would result in the robot falling from the table must be prevented, however further valid movement commands must still
  be allowed.

Create an application that can read in commands of the following form:

```plain
PLACE X,Y,F
MOVE
LEFT
RIGHT
REPORT
MOVE
REPORT
```

- PLACE will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.
- The origin (0,0) can be considered to be the SOUTH WEST most corner.
- The first valid command to the robot is a PLACE command, after that, any sequence of commands may be issued, in any order, including another PLACE command. The application should discard all commands in the sequence until a valid PLACE command has been executed.
- MOVE will move the toy robot one unit forward in the direction it is currently facing.
- LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing the position of the robot.
- REPORT will announce the X,Y and orientation of the robot.
- A robot that is not on the table can choose to ignore the MOVE, LEFT, RIGHT and REPORT commands.
- The application must not exit after the first REPORT command, i.e. many REPORT commands can be received per session.
- It is up to you how you exit the application (e.g. exit command, Ctrl-C etc.)
- Provide test data to exercise the application.

## Constraints

The toy robot must not fall off the table during movement. This also includes the initial placement of the toy robot.
Any move that would cause the robot to fall must be ignored.

Example Input and Output:

```plain
PLACE 0,0,NORTH
MOVE
REPORT
Output: 0,1,NORTH
```

```plain
PLACE 0,0,NORTH
LEFT
REPORT
Output: 0,0,WEST
```

```plain
PLACE 1,2,EAST
MOVE
MOVE
LEFT
MOVE
REPORT
Output: 3,3,NORTH
```

## Project Structure
```markdown
.
├── build             // dockerfile
├── cmd               // command list of CLI
├── config            // command configuration
├── internal          // internal package for robot app
│   ├── errors        // custom errors handler
│   ├── generator     // generate command from template
│   └── parser        // parse command from user input
├── pkg               // package folder of robot app
│   ├── command       // command logic
│   ├── position      // position and direction
│   ├── robot         // robot
│   └── table         // table
└── test              // external test data and helper file
```

## Implementation
The project not only addresses the foundational requirements of the toy robot challenge but also introduces several enhancements to improve the user experience and the system's extensibility:

- Implementation of all basic commands.
- Provision of user-friendly error messages, coupled with a debug mode for detailed error insights.
- Capability to introduce additional commands on-the-fly.
- Flexibility to dynamically set the table size.

### Questions
Given that commands can be case-sensitive and considering the intended audience for this app, it's beneficial to offer a user-friendly interface with permissive input to foster a pleasant experience. 

However, every decision has its trade-offs. From a developer's perspective, it's essential to ensure that all case combinations are comprehended and processed correctly.

Overall, prioritizing a superior user experience aligns with the primary objective for this application.

## Design
When beginning the design of a solution, it's crucial to ensure the flexibility and maintainability of the code. While perfection might not be achieved on the first try, adhering to standard coding practices can aid in making improvements more seamlessly.

Here are some key takeaways from the implementation of this project:
- A Robot is an actor with properties that can be influenced by external actions.
- These external actions, such as commands, should be designed flexibly. They should be implementable without altering the core Robot codebase.
- The Robot doesn't need to understand the intricacies of an action; its primary role is to execute it.
- Drawing a parallel, consider a car. Various cars have distinct designs, but as a driver, your main goal is to press the pedal to set it in motion, without needing a deep understanding of its mechanics. :)
- The table also has a public property for its size. This allows it to be initialized on-the-fly via the CLI, enabling games to start with custom sizes rather than just the hardcoded 5 x 5 dimension.
- While the app will skip any invalid commands as per the requirements, it should still provide a user-friendly message to inform the user about the situation.

## Getting Started

### Prerequisites

- Go 1.20 or higher: [Official installation guide](https://go.dev/doc/install)
- Docker (optional for containerization): [Download here](https://www.docker.com)

### Running the Project
The project can be executed either locally or using Docker.

#### Local

1. Make sure you have all the prerequisites installed.
2. Intall all go packages:
```bash
go mod download
go mod verify
```
3. Run direct with `go` command or via a build
```bash
go run . play
```
or
```bash
go build -o toy-robot
./toy-robot play
```

#### Docker

1. Build image
```bash
@docker build -f ./build/Dockerfile -t toyrobot-app .
```
2. Run container from image
```bash
@docker run -it --rm toyrobot-app play
```

#### Makefile
I also created a Makefile so it can be simple to run via make command
For Docker:
```bash
make build-docker
make run-docker
```
Run directly on your machine:
```bash
make run
```

### Running the tests
```bash
go test ./... -v
```
Run by Makefile
```bash
make test
```
The `test` folder provides a JSON file that can be used to set up extensive test cases and their expected outputs.

You can add additional test cases using the following format to use in `cmd/cmd_test.go`. For example:
```json
{
  "commands": ["PLACE 5,5,NORTH", "REPORT"],
  "output": [
    "Command 'PLACE': invalid position",
    "Command 'REPORT': please place the robot first"
  ]
}
```

### Extra features

#### Running from file
You can place all the commands in a text file and execute them directly without manual typing. Use the following command:
```bash
go run . play -f <your_file_location>
```

The project root also contains a sample command file that you can test:
```bash
go run . play -f ./sample_command.txt
```

#### Running with different table size
You have the option to play the game on a larger board. Use the following command to specify the board's dimensions:

```bash
go run . play --width 10 --height 10
```

These flags can also be combined with the command to run from a file:

```bash
go run . play -f ./sample_command.txt --width 10 --height 10
```

#### Adding a new command template
if you wish to develop another command, the CLI offers an option to generate a template file for a new command. This allows you to easily insert your own logic.
```bash
go run . add command <your_command_name>
```

For example:
```bash
go run . add command jump
```
After running this command, it will generate two new files:
- `pkg/command/jump.go`: jump command logic template
```go
package command

import (
	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
  "github.com/LightningDev/toy-robot-challenge/pkg/table"
)

type JumpCommand struct {
	Name string
}

func NewJumpCommand(args []string) (robot.Command, error) {
	return JumpCommand{
		Name: "JUMP",
	}, nil
}

func (c JumpCommand) Execute(r *robot.Robot, t table.Table) error {
	// TODO: Implement JumpCommand Logic Here
	return nil
}

func (c JumpCommand) GetName() string {
	return c.Name
}

```
- `pkg/command/jump_test.go`: test template for jump command

Furthermore, it updates `pkg/command/command.go` and `config/command.json` to register your new command. These files can be understood as the source of truth, allowing us to check the list of available commands in the app.

`pkg/command/command.go`
```go
var CommandList = map[string]func([]string) (robot.Command, error){
	"PLACE": NewPlaceCommand,
	"REPORT": NewReportCommand,
	"MOVE": NewMoveCommand,
	"LEFT": NewLeftCommand,
	"RIGHT": NewRightCommand,
	"JUMP": NewJumpCommand,
}
```

`config/command.json`
```json
{
  "Commands": [
    "PLACE",
    "REPORT",
    "MOVE",
    "LEFT",
    "RIGHT",
    "JUMP"
  ]
}
```

Afterwards, you can begin implementing your own command logic within the `Execute` function, which is derived from the `Command` interface in `robot`.

#### Running with debug
By default, the app displays user-friendly error messages. However, by using the `-d` flag when running the command, it will also print out both the message and stack trace, enhancing the debugging experience.

```bash
go run . play -d
```