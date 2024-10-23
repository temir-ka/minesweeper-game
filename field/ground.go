package field

// Verify checks if the given cell (row, col) is valid and not a bomb.
func verify(row int, col int, height int, width int, ground [][]int) bool {
	if row < 0 {
		return false
	}
	if row >= height {
		return false
	}
	if col < 0 {
		return false
	}
	if col >= width {
		return false
	}
	if ground[row][col] == -1 {
		return false
	}
	return true
}

func Ground(height int, width int, box [][]rune) [][]int {
	// Create 2D slice
	ground := make([][]int, height)
	for row := 0; row < height; row++ {
		ground[row] = make([]int, width)
		for col := 0; col < width; col++ {
			if box[row][col] == '*' { // If there is a bomb
				ground[row][col] = -1 // We mark bombs
			}
		}
	}
	// Count the adjacent bombs for each cell
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if ground[row][col] == -1 { // If cell is a bomb
				// Increment adjacent cells
				if verify(row-1, col, height, width, ground) {
					ground[row-1][col]++ // Up
				}
				if verify(row+1, col, height, width, ground) {
					ground[row+1][col]++ // Down
				}
				if verify(row, col-1, height, width, ground) {
					ground[row][col-1]++ // Left
				}
				if verify(row, col+1, height, width, ground) {
					ground[row][col+1]++ // Rigt
				}
				if verify(row-1, col-1, height, width, ground) {
					ground[row-1][col-1]++ // Up-Left
				}
				if verify(row-1, col+1, height, width, ground) {
					ground[row-1][col+1]++ // Up-Right
				}
				if verify(row+1, col-1, height, width, ground) {
					ground[row+1][col-1]++ // Down-Left
				}
				if verify(row+1, col+1, height, width, ground) {
					ground[row+1][col+1]++ // Down-Right
				}
			}
		}
	}
	return ground
}
