package field

import "math/rand"

// change box when first move was in bomb cell. Shift bomb cell on random empty cell without bomb.
func ChangeBox(row int, col int, height int, width int, box [][]rune) [][]rune {
	for true {
		new_row := rand.Intn(height)
		new_col := rand.Intn(width)
		if box[new_row][new_col] == '*' {
			continue
		} else {
			box[new_row][new_col] = '*'
			break
		}
	}
	box[row][col] = '.'
	return box
}
