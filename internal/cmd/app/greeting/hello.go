package greeting

import (
	"fmt"
)

type MyEvent struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (m MyEvent) SayHello() {
	fmt.Printf("Hello %s: %s", m.Name, m.Message)
}
