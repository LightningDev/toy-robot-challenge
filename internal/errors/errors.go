package errors

import (
	"fmt"
	"log"
)

type ValidationError struct {
	Command string
	Err     error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Command, e.Err.Error())
}

func HandleError(err error) {
	switch e := err.(type) {
	case *ValidationError:
		log.Printf("Validation error in field %s: %s", e.Command, e.Err)
	default:
		log.Printf("Unknown error: %s", err)
	}
}
