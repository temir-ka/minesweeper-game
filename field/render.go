package field

import (
	"github.com/alem-platform/ap"
)

func Render(height int, width int, ground [][]int, check [][]bool) {
	Col_Digits(width) // Display column digits
	Roof(width)       // Display the roof of the map

	for row := 0; row < height; row++ {
		for stage := 0; stage < 3; stage++ {
			Row_Digits(stage, row+1) // Display row digits
			for col := 0; col < width; col++ {
				if !check[row][col] { // If cell is not opened
					Close_Cell(9)
				} else { // If cell is open
					if ground[row][col] == -1 { // If opened cell is bomb
						Bomb(stage, ground[row][col])
					} else {
						Open_Cell(stage, ground[row][col])
					}
				}
			}
			ap.PutRune('|')
			ap.PutRune('\n')
		}
	}
}
