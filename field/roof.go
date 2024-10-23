package field

import "github.com/alem-platform/ap"

// print roof
func Roof(width int) {
	// Spaces
	for i := 0; i < 4; i++ {
		ap.PutRune(' ')
	}
	for i := 0; i < 7*width+width-1; i++ {
		ap.PutRune('_')
	}
	ap.PutRune(' ')
	ap.PutRune('\n')
}
