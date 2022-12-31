package gotermio

import "fmt"

type Style  int
type Styles []Style
const (
	Bold      = Style(1)
	Italics   = Style(3)
	Underline = Style(4)
	Blink     = Style(5)
	Invert    = Style(7)
	Strike    = Style(9)
)

func StyleString(section string, styles Styles) (s string) {
	for _, style := range styles {
		s += fmt.Sprintf("\x1b[%vm", int(style))
	}

	return fmt.Sprintf("%v%v\x1b[0m", s, section)
}

func SetStyle(styles Styles) {
	for _, style := range styles {
		fmt.Printf("\x1b[%vm", int(style))
	}
}

func ResetStyle() {
	fmt.Printf("\x1b[0m")
}
