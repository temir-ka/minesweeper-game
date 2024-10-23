package field

import "github.com/alem-platform/ap"

// print simple string
func PrintString(str string, new_line bool) {
	for i := range str {
		ap.PutRune(rune(str[i]))
	}
	if new_line {
		ap.PutRune('\n')
	}
}

// print grid size for end statistics
func PrintFieldSize(height int, width int) {
	PrintString("- Field size: ", false)
	PutNumber(height)
	ap.PutRune('x')
	PutNumber(width)
	ap.PutRune('\n')
}

// print string with int number
func PrintStringNumber(str string, number int) {
	PrintString(str, false)
	PutNumber(number)
	ap.PutRune('\n')
}
