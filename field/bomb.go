package field

import "github.com/alem-platform/ap"

// Bomb displays a bomb cell
func Bomb(stage int, n int) {
	ap.PutRune('|')
	Color(n) // Sets the color based on the cell value (n)
	for i := 0; i < 7; i++ {
		if stage == 0 {
			ap.PutRune(' ')
		} else if stage == 1 {
			if i == 3 {
				ap.PutRune('*') //  Display a '*' to indicate the bomb in the center
			} else {
				ap.PutRune(' ')
			}
		} else {
			ap.PutRune('_') // If stage == 3, print only underscores
		}
	}
	Reset() // Resets the color
}

func BombCount(height int, width int, ground [][]int) int {
	bombs := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if ground[i][j] == -1 {
				bombs++
			}
		}
	}
	return bombs
}
