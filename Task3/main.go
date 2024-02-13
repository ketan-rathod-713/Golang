package main

import (
	"fmt"
	"log"
)

type SudokuData struct {
	Sudoku     [][]int
	IsValid    bool
	IsSolvable bool
}

var Sudokus []SudokuData = []SudokuData{
	{
		Sudoku: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		IsValid:    true,
		IsSolvable: true,
	},
	{
		Sudoku: [][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 4, 5, 2, 8, 6, 1, 7, 9},
		},
		IsValid:    true,
		IsSolvable: true,
	},
	{
		Sudoku: [][]int{
			{0, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 0, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 0, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 0, 1, 5, 3, 7, 2, 8, 4},
			{2, 0, 7, 4, 1, 9, 6, 3, 5},
			{3, 4, 5, 2, 8, 6, 1, 7, 9},
		},
		IsValid:    true,
		IsSolvable: true,
	},
	{
		Sudoku: [][]int{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		},
		IsValid:    true,
		IsSolvable: true,
	},
	{
		Sudoku: [][]int{
			{5, 3, 4, 6, 7, 8, 2, 1, 2},
			{6, 7, 2, 1, 0, 5, 0, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 4, 5, 2, 8, 6, 1, 7, 9},
		},
		IsValid:    false,
		IsSolvable: false,
	},
	{
		Sudoku: [][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 0, 4, 1, 9, 6, 7, 5},
			{3, 4, 5, 2, 8, 6, 1, 0, 9},
		},
		IsValid:    true,
		IsSolvable: false,
	},
}

func main() {
	log.Println("Welcome To Sudoku Solver")
	/* Test All Sudoku */
	for _, sudoku := range Sudokus {
		validityOfSudoku, row, col := CheckValidSudoku(sudoku.Sudoku)
		if validityOfSudoku {
			ans := CheckSudokuSolvable(&sudoku.Sudoku, 0, 0)
			if ans {
				log.Println("Below Sudoku is Solvable")
			} else {
				log.Println("Below Sudoku is Not Solvable")
			}
			printSudoku(sudoku.Sudoku)
		} else {
			log.Println("Invalid Sudoku Provided, See Below")
			log.Printf("It Produces Error On Row %v and Column %v \n", row, col)
			printSudoku(sudoku.Sudoku)
		}
	}
}

/*
Checks if sudoku is valid or not
Using Set
Returns (isValid, Row, Column)
Row == -1 and Column == -1 if valid sudoku
else it will point to the cell which producing error
*/
func CheckValidSudoku(sudoku [][]int) (bool, int, int) {
	// check all 3 condtions
	mp := make(map[int]bool)

	// Row
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			_, exist := mp[sudoku[i][j]]
			if exist { // if already exist then return false as we want set here
				return false, i, j
			} else {
				if sudoku[i][j] != 0 {
					mp[sudoku[i][j]] = true

				}
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
				return false, j, i
			} else {
				if sudoku[j][i] != 0 {
					mp[sudoku[j][i]] = true
				}
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
						return false, m, n
					} else {
						if sudoku[m][n] != 0 {
							mp[sudoku[m][n]] = true
						}
					}
				}
			}

			// clear map
			for k := range mp {
				delete(mp, k)
			}

		}
	}

	return true, -1, -1 // no error
}

/*
Checks if valid sudoku is solvable or not
Using backtracking
*/
func CheckSudokuSolvable(sudoku *[][]int, i int, j int) bool {

	// If sudoku is not valid then it is not solvable
	ok, _, _ := CheckValidSudoku(*sudoku)
	if ok == false {
		return false
	}

	// Get Next Sudoku Cell
	nextCellI, nextCellJ := nextCell(i, j)

	var ans bool = false

	/*	If Empty Cell
		Then : Try Values From 1 to 9 inclusive AND Check If Valid Sudoku can be Produced From it or not
		Else : Already Placed Value -> Ignore it and move on to check next cell

		If Any Try Value Results in valid sudoku then immediatly return true ( as we want only one possible solution )
	*/
	if (*sudoku)[i][j] == 0 { // given sudoku cell has empty value, i can try inserting 1 to 9 // let's see which works
		// Check If I can insert value from 1 to 9

		for val := 1; val <= 9; val++ {
			// Check if Value Can be inserted in cell or not // If Yes then insert value
			if checkIsValidToInsertCellValue(*sudoku, i, j, val) {
				(*sudoku)[i][j] = val

				// If last cell THEN sudoku is solved // return true
				if isLastCell(i, j) {
					return true
				}

				// If not last cell THEN check for next remaining cells
				ans = CheckSudokuSolvable(sudoku, nextCellI, nextCellJ)
				if ans == true { // Return Here Only As Now We don't want to try other combinations
					return true
				} else {
					// backtrack if this trial value found producing unsolvable sudoku (eg. we got false value )
					(*sudoku)[i][j] = 0
				}
			} else {
				continue
			}
		}
	} else {

		// For Already Placed Values
		if isLastCell(i, j) {
			return true
		}

		// If not last cell THEN check for next remaining cells
		ans = CheckSudokuSolvable(sudoku, nextCellI, nextCellJ)
	}

	return ans
}

// TODO : do unit testing using testing package
func checkIsValidToInsertCellValue(sudoku [][]int, i int, j int, val int) bool {

	// Check If i can insert value in given row
	for it := 0; it < 9; it++ {
		if sudoku[i][it] == val {
			return false
		}
	}

	// Check If i can insert value in given column
	for it := 0; it < 9; it++ {
		if sudoku[it][j] == val {
			return false
		}
	}

	/*
		Total Boxes in sudoku : 9 of size 3*3
		using boxI and boxJ we can get combinations like 00, 01,... 22 => giving 9 boxes.
	*/
	boxI := i / 3
	boxJ := j / 3

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

// check if last cell of sudoku
func isLastCell(i int, j int) bool {
	if i == 8 && j == 8 {
		return true
	}

	return false
}

// Print sudoku
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
