package errs

import (
	"fmt"
)

type EventNotFound struct{}

func (e *EventNotFound) Error() string {
	return "Event record not found."
}

type ReadEventException struct {
	Err error
}

func (e *ReadEventException) Error() string {
	return fmt.Sprintf("Failed to read event. [error]: %s", e.Err)
}
