package field

import "math/rand"

func GenerateRandomGrid(height, width int) [][]rune {
	totalCells := height * width
	// Bomb count (15% of cells)
	numBombs := totalCells * 15 / 100
	// Minimum 2 bombs
	if numBombs < 2 {
		numBombs = 2
	}
	// Empty grid
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	// Insert bombs
	for numBombs > 0 {
		row := rand.Intn(height)
		col := rand.Intn(width)
		if grid[row][col] != '*' {
			grid[row][col] = '*'
			numBombs--
		}
	}
	return grid
}
