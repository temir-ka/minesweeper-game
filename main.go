package main

import (
	"crunch02/field"
	"fmt"
)

func main() {
	field.PrintString("Choose a mode: ", true)
	field.PrintString("1. Enter a custom map", true)
	field.PrintString("2. Generate a random map", true)
	// process mode choice
	field.PrintString("Enter your choice: ", false)
	var ch string
	fmt.Scanf("%s ", &ch)
	choice := field.ChoiceError(ch)
	if choice == -1 {
		field.PrintString("Invalid input.", true)
		return
	}
	// process grid size
	field.PrintString("Enter the size of the grid: ", false)
	var h, w string
	fmt.Scanf("%s %s ", &h, &w)
	height := field.HWError(h)
	width := field.HWError(w)
	if height == -1 || width == -1 {
		field.PrintString("Invalid input.", true)
		return
	}
	// creating "box" array with "." (empty) and "*" (bomb)
	// creating "check" array for following player movement
	box := make([][]rune, height)
	check := make([][]bool, height)
	for i := 0; i < height; i++ {
		box[i] = make([]rune, width)
		check[i] = make([]bool, width)
	}
	// custom or random grid
	if choice == 1 {
		for i := 0; i < height; i++ {
			var str string
			fmt.Scanf("%s ", &str)
			if field.FieldError(str, width) == false {
				field.PrintString("Invalid input.", true)
				return
			}
			for j := 0; j < width; j++ {
				box[i][j] = rune(str[j])
			}
		}
	} else {
		box = field.GenerateRandomGrid(height, width)
	}
	// creating and computing array with adjacent count 
	ground := make([][]int, height)
	ground = field.Ground(height, width, box)

	bombs := field.BombCount(height, width, ground)
	if bombs < 2 {
		field.PrintString("There must be at least two bombs", true)
		return
	}

	// print grid
	field.Render(height, width, ground, check)
	game := true
	moves := 0
	first_move := true
	// start game
	for game {
		field.PrintString("Enter coordinates: ", false)
		var r, c string
		fmt.Scanf("%s %s ", &r, &c)
		// process error inputs
		row := field.RCError(r)
		col := field.RCError(c)
		if row == -1 || col == -1 {
			field.PrintString("Invalid input.", true)
			continue
		}
		row--
		col--
		if row >= height || col >= width {
			field.PrintString("Invalid input.", true)
			continue
		}
		moves++
		// if first move, we change bomb location
		if first_move {
			if ground[row][col] == -1 {
				box = field.ChangeBox(row, col, height, width, box)
				ground = field.Ground(height, width, box)
			}
			first_move = false
		}
		// check player's selected cell
		if ground[row][col] == -1 {
			field.Fail(height, width, ground, check)
			game = false
		} else if ground[row][col] == 0 {
			field.AdjacentOpen(row, col, height, width, ground, check)
		} else {
			check[row][col] = true
		}
		field.Render(height, width, ground, check)
		// check game end
		if game == false {
			field.PrintString("Game Over!", true)
			field.PrintString("Your Statistics:", true)
			field.PrintFieldSize(height, width)
			field.PrintStringNumber("- Number of bombs: ", bombs)
			field.PrintStringNumber("- Number of moves: ", moves)
		}
		if field.CheckWin(check, bombs) {
			game = false
			field.PrintString("You Win!", true)
			field.PrintString("Your Statistics:", true)
			field.PrintFieldSize(height, width)
			field.PrintStringNumber("- Number of bombs: ", bombs)
			field.PrintStringNumber("- Number of moves: ", moves)
		}
	}
}
