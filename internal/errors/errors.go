package errors

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

var Debug bool

type ValidationError struct {
	Command string
	Err     error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Command, e.Err.Error())
}

func HandleError(err error) {
	if Debug {
		stack := debug.Stack()
		log.Printf("[DEBUG] Stack trace:\n%s", stack)
	} else {
		log.SetFlags(0)
	}

	switch e := err.(type) {
	case *ValidationError:
		if Debug {
			log.Printf("Validation error in field %s: %s", e.Command, e.Err)
		} else {
			log.Printf("Issue with the command '%s': %s", e.Command, e.Err)
		}
	default:
		if Debug {
			log.Printf("Unknown error: %+v", err)
		} else {
			log.Println("Sorry, an unexpected error occurred. Please try again!")
		}
	}
}

func New(message string) error {
	return errors.New(message)
}
