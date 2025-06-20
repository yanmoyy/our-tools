package color

type Color int

const (
	Red Color = iota
	Green
	Yellow
	Blue
)

const (
	reset = "\033[0m"
)

func getANSIColor(color Color) string {
	switch color {
	case Red:
		return "\033[31m"
	case Green:
		return "\033[32m"
	case Yellow:
		return "\033[33m"
	case Blue:
		return "\033[34m"
	default:
		return ""
	}
}

func (color Color) ColorString(s string) string {
	return getANSIColor(color) + s + reset
}
