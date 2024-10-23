package field

import "github.com/alem-platform/ap"

// Color function sets the terminal color based on the provided integer n.
func Color(n int) {
	var color []rune
	// Select the ANSI color code based on the value of n
	switch n {
	case 0:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '3', '7', 'm'}
	case -1:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '1', '9', '6', 'm'}
	case 1:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '2', '8', 'm'}
	case 2:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '2', '2', '0', 'm'}
	case 3:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '2', '1', '3', 'm'}
	case 4:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '1', '9', 'm'}
	case 5:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '1', '7', '2', 'm'}
	case 6:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '2', '1', '8', 'm'}
	case 7:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '1', '4', '6', 'm'}
	case 8:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '1', '0', '6', 'm'}
	case 9:
		color = []rune{'\033', '[', '4', '8', ';', '5', ';', '2', '3', '9', 'm'}
	}
	for _, char := range color {
		ap.PutRune(char)
	}
}

// Reset function resets the terminal color to the default.
func Reset() {
	reset := []rune{'\033', '[', '0', 'm'} // ANSI code to reset the color
	for _, char := range reset {
		ap.PutRune(char)
	}
}
