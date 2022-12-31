package main

import (
	"fmt"

	term "github.com/LordOfTrident/gotermio"
)

var (
	h = 0

	notifStyle = term.Styles{term.Bold, term.Style(term.Bright(term.White))}
)

func cbreak() {
	term.CursorGoto(0, h - 1)

	fmt.Print(term.StyleString("CTRL+C Pressed!", notifStyle))
	term.ClearFromCursorToLineEnd()

	term.CursorGoto(0, 1)
}

func resized() {
	_, h = term.GetSize()

	term.CursorGoto(0, h - 1)

	fmt.Print(term.StyleString("Resized!", notifStyle))
	term.ClearFromCursorToLineEnd()

	term.CursorGoto(0, 1)
}

func main() {
	_, h = term.GetSize()

	term.CBreak(cbreak)
	term.Resized(resized)

	term.HideCursor()

	fmt.Println(term.StyleString("Started screen test", term.Styles{term.Bold}))

	term.AltScreen(true)
	term.CursorGoto(0, 0)

	fmt.Print("Hello there!")

	term.GetKey()

	term.CursorUp(1)
	term.CursorRight(2)
	term.ClearFromCursorToLineEnd()

	fmt.Print("Bye")

	term.GetKey()
	term.AltScreen(false)

	fmt.Println(term.StyleString("Ended screen test", term.Styles{term.Invert}))

	term.ShowCursor()
}
