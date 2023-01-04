package model

import (
	"fmt"
)

type Greeting struct {
	UserName string `json:"name"`
	Message  string `json:"message"`
}

func (g *Greeting) Hello() string {
	return fmt.Sprintf("%s: %s", g.UserName, g.Message)
}
