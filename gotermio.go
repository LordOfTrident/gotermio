package gotermio

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"os/signal"
	"syscall"
	"regexp"
)

const (
	Icanon = iota << 0
	echo   = iota << 1
)

type Settings struct {
	icanon bool
	echo   bool
}

// 'stty size' output format is '<HEIGHT> <WIDTH>'
var sizeRegex = regexp.MustCompile("([0-9]*)\\s([0-9]*)\\s*")

func GetSize() (int, int) {
	bytes, err := exec.Command("stty", "-F", "/dev/tty", "size").Output()
	if err != nil {
		panic(err)
	}

	// Match the output format
	info := sizeRegex.FindStringSubmatch(string(bytes))

	var h, w int

	// Parse the strings
	h, err = strconv.Atoi(info[1])
	if err != nil {
		panic(err)
	}

	w, err = strconv.Atoi(info[2])
	if err != nil {
		panic(err)
	}

	return w, h
}

func Resized(resized func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGWINCH)

	go func() {
		for {
			<- c
			resized()
		}
	}()
}

func CBreak(cbreak func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			<- c
			cbreak()
		}
	}()
}

func GetKey() byte {
	in := make([]byte, 1)
	os.Stdin.Read(in)

	return in[0]
}

func AltScreen(enable bool) {
	if enable {
		fmt.Print("\x1b[?1049h")
	} else {
		fmt.Print("\x1b[?1049l")
	}
}

func ClearFromCursorToScreenEnd() {
	fmt.Print("\x1b[0J")
}

func ClearFromCursorToScreenStart() {
	fmt.Print("\x1b[1J")
}

func ClearScreen() {
	fmt.Print("\x1b[2J")
}

func ClearFromCursorToLineEnd() {
	fmt.Print("\x1b[0K")
}

func ClearFromCursorToLineStart() {
	fmt.Print("\x1b[1K")
}

func ClearLine() {
	fmt.Print("\x1b[2K")
}
