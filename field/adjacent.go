package field

// AdjacentOpen recursively reveals adjacent cells in the Minesweeper map
func AdjacentOpen(row int, col int, height int, width int, ground [][]int, check [][]bool) {
	// Return if the cell is out of bounds, is a bomb, or has already been revealed
	if row < 0 || row >= height || col < 0 || col >= width || ground[row][col] == -1 || check[row][col] == true {
		return
	}
	check[row][col] = true // Mark the current cell as opened
	if ground[row][col] > 0 { // If cell is not '0', stop the function
		return
	}
	// Open all adjacent cells
	AdjacentOpen(row-1, col, height, width, ground, check)   // Up
	AdjacentOpen(row+1, col, height, width, ground, check)   // Down
	AdjacentOpen(row, col-1, height, width, ground, check)   // Left
	AdjacentOpen(row, col+1, height, width, ground, check)   // Right
	AdjacentOpen(row-1, col-1, height, width, ground, check) // Up-Left
	AdjacentOpen(row-1, col+1, height, width, ground, check) // Up-Right
	AdjacentOpen(row+1, col-1, height, width, ground, check) // Down-Left
	AdjacentOpen(row+1, col+1, height, width, ground, check) // Down-Right
}
