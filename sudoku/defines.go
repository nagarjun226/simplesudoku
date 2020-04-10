package simplesudoku

import "fmt"

/*
Since golang uses a list of lists to support matrices while making operations on the table, we must remember
the row column notation as the first for loop will extract the row and the second essentially extracts the
element from the row so a Sudoku is [row][column]. Basically, we are using the (row, column) fomrat and not the
(i,j) format

In this code the row and columns are indexed from 0

Zerospaces in our Sudoku are treated as empty spaces
*/

// SudokuT - Define a Sudoku as a 9x9 grid of Ints
// Note: Technically a sudoku has only 9 values so using int may be quite wasteful
type SudokuT [9][9]int

// GetSudoku - Constructor for a Sudoku object
func GetSudoku() *SudokuT {
	s := new(SudokuT)
	return s
}

// Print - Print a Sudoku in Human Readable Format
func (s *SudokuT) Print() {
	for _, row := range s {
		for _, element := range row {
			fmt.Printf("%d\t", element)
		}
		fmt.Println()
	}
}

// GetElement - return the element at (row, column)
func (s *SudokuT) GetElement(row, column int) (int, error) {
	if row > 8 || column > 8 || row < 0 || column < 0 {
		return 0, fmt.Errorf("Invalid Inputs")
	}
	n := s[row][column]
	return n, nil
}

// SetElement - set the element at (row, column)
func (s *SudokuT) SetElement(row, column, n int) error {
	if row > 8 || column > 8 || row < 0 || column < 0 || n > 9 || n < 0 {
		return fmt.Errorf("Invalid Inputs")
	}
	s[row][column] = n
	return nil
}

// SetFromArray - Set the elements of a Sudoku from a array of length 81
// Sudoku [r1;r2;...;r9] = [r1,r2,...,r9]
func (s *SudokuT) SetFromArray(l []int) error {
	if len(l) != 81 {
		return fmt.Errorf("Array length not 81")
	}

	index := 0

	for r := range s {
		for c := range s[r] {
			err := s.SetElement(r, c, l[index])
			if err != nil {
				return err
			}
			index++
		}
	}

	if index != 80 {
		return fmt.Errorf("Something Went Wrong setting the sudoku from array")
	}

	return nil

}

// ArrayFormat - Represent the elements of a Sudoku as an array of length 81
// Sudoku [r1,r2,...,r9] = [r1;r2;...;r9]
func (s *SudokuT) ArrayFormat() ([]int, error) {
	a := make([]int, 0)
	for r := range s {
		for c := range s[r] {
			e, err := s.GetElement(r, c)
			if err != nil {
				return []int{}, err
			}
			a = append(a, e)
		}
	}

	if len(a) != 81 {
		return a, fmt.Errorf("Lenth of the array not 81 len = %v", len(a))
	}

	return a, nil
}
