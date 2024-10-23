package field

// check game win or yet not
func CheckWin(check [][]bool, bombs int) bool {
	// count of closed cells
	closed_count := 0
	for i := range check {
		for j := range check[i] {
			if check[i][j] == false {
				closed_count += 1
			}
		}
	}
	// if closed cells count equals bombs count then game was won
	if closed_count == bombs {
		return true
	}
	return false
}
