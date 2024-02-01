package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome To Sudoku Solver \n")

	// input := [][]int{
	// 	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	// 	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	// 	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	// 	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	// 	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	// 	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	// 	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	// 	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	// 	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	// }

	// unsolvable_sudoku := [][]int{
	// 	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	// 	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	// 	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	// 	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	// 	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	// 	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	// 	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	// 	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	// 	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	// }

	exampleInput := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	unsolvable_sudoku := exampleInput
	// emptySudoku := [][]int{
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }

	// printSudoku(emptySudoku)

	ans := checkSudokuSolvable(&unsolvable_sudoku, 0, 0)

	if ans {
		fmt.Println("Sudoku is solvable")
		printSudoku(unsolvable_sudoku)
	} else {
		fmt.Println("Sudoku is not solvable")
		printSudoku(unsolvable_sudoku)
	}

	ans2 := checkValidSudoku(unsolvable_sudoku)
	fmt.Println(ans2)
}

/*
	Checks if sudoku is valid or not
*/
func checkValidSudoku(sudoku [][]int) bool {
	// check all 3 condtions
	mp := make(map[int]bool)

	// Row
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			_, exist := mp[sudoku[i][j]]
			if exist { // if already exist then return false as we want set here
				return false
			} else {
				mp[sudoku[i][j]] = true
			}
		}
		for k := range mp {
			delete(mp, k)
		}
	}

	// Column
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			_, exist := mp[sudoku[j][i]]
			if exist { // if already exist then return false as we want set here
				return false
			} else {
				mp[sudoku[j][i]] = true
			}
		}
		for k := range mp {
			delete(mp, k)
		}
	}

	// Boxes of 3*3
	// For all boxes 2 loops == 3*3 == 9 boxes
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			//Now we are inside a box //  For Each cell // 3*3 // 9 cells for each box
			for m := i * 3; m < i*3+3; m++ {
				for n := j * 3; n < j*3+3; n++ {
					_, exist := mp[sudoku[m][n]]
					if exist { // if already exist then return false as we want set here
						return false
					} else {
						mp[sudoku[m][n]] = true
					}
				}
			}

			// clear map
			for k := range mp {
				delete(mp, k)
			}

		}
	}

	return true
}

/*
	Checks if valid sudoku is solvable or not
	- Using backtracking
*/
func checkSudokuSolvable(sudoku *[][]int, i int, j int) bool {
	// fmt.Println("Check Sudoku For ", i, j)
	nextCellI, nextCellJ := nextCell(i, j)

	var ans bool = false
	if (*sudoku)[i][j] == 0 { // given sudoku cell has empty value, i can try inserting 1 to 9 // let's see which works
		// Check If I can insert value from 1 to 9

		for val := 1; val <= 9; val++ {
			// check if valid to insert for given sudoku

			if checkIsValidToInsertCellValue(*sudoku, i, j, val) {
				(*sudoku)[i][j] = val

				if isLastCell(i, j) {
					printSudoku(*sudoku)
					return true
				}

				ans = checkSudokuSolvable(sudoku, nextCellI, nextCellJ)
				if ans == true { // Return Here Only As Now We don't want to try other combinations
					return true
				} else {
					// if ans is not true then mark given location with 0
					(*sudoku)[i][j] = 0
				}

			} else {
				continue
			}
		}
	} else { // for Already Placed Values
		// If last cell then no need to do anything just return true
		if isLastCell(i, j) {
			return true
		}

		// what if Here I check for one condtion that is if valid sudoku or not
		// if here it is not valid then previous arrangements will be false hence
		// TODO

		//if not last cell then call CSS for next cell
		ans = checkSudokuSolvable(sudoku, nextCellI, nextCellJ)
	}

	return ans
}

// Check if it is a valid one or not
// TODO : do unit testing if required
func checkIsValidToInsertCellValue(sudoku [][]int, i int, j int, val int) bool {

	for it := 0; it < 9; it++ { // for all 9 rows
		if sudoku[i][it] == val {
			return false
		}
	}

	// Check Column
	for it := 0; it < 9; it++ { // for all 9 rows
		if sudoku[it][j] == val {
			return false
		}
	}

	boxI := i / 3
	boxJ := j / 3 // both can be either 0,1,2 // which makes combinations of 00, 01, 02 ... 9 comb -> 9 boxes

	for it := boxI * 3; it < boxI*3+3; it++ { // outer row
		for jt := boxJ * 3; jt < boxJ*3+3; jt++ {
			if sudoku[it][jt] == val {
				return false
			}
		}
	}

	return true
}

/*
	It returns the next cell of sudoku to search
	for eg. nextCell(1, 8) will return (2, 0) (Assumption : Iterating row by row)
*/
func nextCell(i int, j int) (int, int) {
	if j == 8 {
		return i + 1, 0
	} else {
		return i, j + 1
	}
}

func isLastCell(i int, j int) bool {
	if i == 8 && j == 8 {
		return true
	}

	return false
}

func printSudoku(sudoku [][]int) {

	fmt.Println("")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(sudoku[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
