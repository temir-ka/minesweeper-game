package field

import "github.com/alem-platform/ap"

// creating digit sequence on x-axis
func Col_Digits(width int) {
	var num int = 1 // Start numberinh from 1
	for i := 0; i < width; i++ {
		for j := 0; j < 7-LenNumber(num)+1; j++ {
			ap.PutRune(' ')
		}
		PutNumber(num) // Display the current column number
		num++          // Incrrement the number for the next column
	}
	ap.PutRune('\n')
}

// creating digit sequence on y-axis
func Row_Digits(stage int, number int) {
	var spaces int = 3 // Number of spaces
	if stage == 1 {
		var copy int = number
		// Calculate the number of spaces based on length of number
		for copy > 0 {
			spaces--
			copy /= 10
		}
		PutNumber(number) // Display the current row number
		// Display spaces after the number
		for i := 0; i < spaces; i++ {
			ap.PutRune(' ')
		}
	} else {
		for i := 0; i < spaces; i++ {
			ap.PutRune(' ')
		}
	}
}
