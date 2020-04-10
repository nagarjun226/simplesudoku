package main

import (
	"fmt"

	simplesudoku "github.com/nagarjun226/simplesudoku/sudoku"
)

func main() {
	// playground
	a := []int{0, 3, 5, 2, 6, 9, 7, 8, 1, 6, 8, 2, 5, 7, 1, 4, 9, 3, 1, 9, 7, 8, 3, 4, 5, 6, 2, 8, 2, 6, 1, 9, 5, 3, 4, 7, 3, 7, 4, 6, 8, 2, 9, 1, 5, 9, 5, 1, 7, 4, 3, 6, 2, 8, 5, 1, 9, 3, 2, 6, 8, 7, 4, 2, 4, 8, 9, 5, 7, 1, 3, 6, 7, 6, 3, 4, 1, 8, 2, 5, 9}
	s := simplesudoku.GetSudoku()
	e := s.SetFromArray(a)
	fmt.Println(e)
	s.Print()
	ar, err := s.ArrayFormat()
	fmt.Println("\n", ar, "\n", err)

	v1, err := s.ValidMove(0, 0, 4)
	fmt.Println(v1, err)
}
