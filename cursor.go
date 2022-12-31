package gotermio

import "fmt"

type CursorShape int
const (
	CursorDefault = CursorShape(iota)
	CursorBlockBlink
	CursorBlock
	CursorUnderlineBlink
	CursorUnderline
	CursorBarBlink
	CursorBar
)

func ShapeCursor(shape CursorShape) {
	fmt.Printf("\x1b[%v q", int(shape))
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func CursorGoto(x, y int) {
	fmt.Printf("\x1b[%v;%vH", y + 1, x + 1)
}

func CursorUp(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vB", -by)
	} else {
		fmt.Printf("\x1b[%vA", by)
	}
}

func CursorDown(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vA", -by)
	} else {
		fmt.Printf("\x1b[%vB", by)
	}
}

func CursorRight(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vD", -by)
	} else {
		fmt.Printf("\x1b[%vC", by)
	}
}

func CursorLeft(by int) {
	if by == 0 {
		return
	} else if by < 0 {
		fmt.Printf("\x1b[%vC", -by)
	} else {
		fmt.Printf("\x1b[%vD", by)
	}
}
