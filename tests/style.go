package main

import (
	"fmt"

	term "github.com/LordOfTrident/gotermio"
)

func main() {
	styles := []term.Style{term.Bold, term.Style(term.Red)}

	fmt.Printf("%v, world!\n", term.StyleString("Hello", styles))
}
