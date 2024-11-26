package main

import (
	"testing"
)

func TestCheckValidSudoku(t *testing.T) {

	for _, s := range Sudokus {
		got, _, _ := CheckValidSudoku(s.Sudoku)
		if got == s.IsValid {
			t.Log("Got it correct")
		} else {
			t.Fatal("Want ", s.IsValid, " but got ", got)
		}
	}
}

func TestCheckSudokuSolvable(t *testing.T) {

	for i, s := range Sudokus {
		got := CheckSudokuSolvable(&s.Sudoku, 0, 0)
		if got == s.IsSolvable {
			t.Log("Got it correct")
		} else {
			t.Fatal("Want ", s.IsSolvable, " but got ", got, " for ", i, s)
		}
	}
}
