package project

import (
	"fmt"
)

type Robot struct {
	Name string
	uuid	string
	model	string
}

//go:generate echo hello, world

func main() {
	fmt.Println("Hello world!")
}
