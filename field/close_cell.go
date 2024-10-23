package field

import "github.com/alem-platform/ap"

// Close_Cell displays closed cells
func Close_Cell(n int) {
	ap.PutRune('|')
	Color(n) // Sets the color based on the cell value (n)
	for i := 0; i < 7; i++ {
		ap.PutRune('X')
	}
	Reset() // Resets the color
}
