package gotermio

type Color int
const (
	Black = Color(iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func IsBright(color Color) bool {
	return int(color) >= 90 && int(color) < 110
}

func IsDark(color Color) bool {
	return !IsBright(color)
}

func IsFg(color Color) bool {
	return int(color) / 10 % 3 == 0
}

func IsBg(color Color) bool {
	return !IsFg(color)
}

func Bright(color Color) Color {
	if IsBright(color) {
		return color
	}

	return Color(int(color) + 60)
}

func Dark(color Color) Color {
	if IsDark(color) {
		return color
	}

	return Color(int(color) - 60)
}

func Fg(color Color) Color {
	if IsFg(color) {
		return color
	}

	return Color(int(color) - 10)
}

func Bg(color Color) Color {
	if IsBg(color) {
		return color
	}

	return Color(int(color) + 10)
}
