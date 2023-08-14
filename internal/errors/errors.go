package errors

import (
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
			log.Printf("[DEBUG] Validation error in field %s: %s", e.Command, e.Err)
		} else {
			log.Printf("Command '%s': %s", e.Command, e.Err)
		}
	default:
		if Debug {
			log.Printf("[DEBUG] Unknown error: %+v", err)
		} else {
			log.Println("Sorry, an unexpected error occurred. Please try again!")
		}
	}
}
