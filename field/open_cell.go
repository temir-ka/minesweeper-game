package field

import "github.com/alem-platform/ap"

// Open_Cell displays a open cell
func Open_Cell(stage int, n int) {
	ap.PutRune('|')
	Color(n) // Sets the color based on the cell value (n)
	// Loop to display the contents of the cell
	for i := 0; i < 7; i++ {
		if stage == 0 {
			ap.PutRune(' ')
		} else if stage == 1 {
			if i == 3 {
				if n == 0 {
					ap.PutRune(' ')
				} else {
					ap.PutRune(rune(n + '0'))
				}
			} else {
				ap.PutRune(' ')
			}
		} else {
			ap.PutRune('_') // If stage == 3, print only underscores
		}
	}
	Reset() // Resets the color
}
