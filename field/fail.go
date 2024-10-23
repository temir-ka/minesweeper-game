package field

// open all bomb cells when game was losed.
func Fail(height int, width int, ground [][]int, check [][]bool) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if ground[row][col] == -1 {
				check[row][col] = true
			}
		}
	}
}
