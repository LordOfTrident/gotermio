package main

import (
	"fmt"

	term "github.com/LordOfTrident/gotermio"
)

func main() {
	fmt.Println("Hello, world!")
	term.CursorUp(1)
	fmt.Println("Hi")
}
