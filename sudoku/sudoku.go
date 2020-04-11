package simplesudoku

import "fmt"

// ValidMove - input parameters is row, column and a number between 1 and 9
// The function evaluates if placing the number n in the cell (row, column)
// is a possible move or not depending on the present state of Sudoku s
// Note: Multiple n's may be possible for a s[row][column]
func (s *SudokuT) ValidMove(row, column, n int) (bool, error) {
	if row > 8 || column > 8 || row < 0 || column < 0 || n > 9 || n < 1 {
		return false, fmt.Errorf("Invalid Inputs")
	}

	if s[row][column] != 0 {
		return false, fmt.Errorf("Index already has a value")
	}

	// Check row validity
	var colMatch bool
	for ii := 0; ii < 9; ii++ {
		if ii == column {
			continue
		}
		if s[row][ii] == n {
			colMatch = true
			break
		}
	}

	if colMatch {
		return false, nil
	}

	// Check column validity
	var rowMatch bool
	for ii := 0; ii < 9; ii++ {
		if ii == row {
			continue
		}
		if s[ii][column] == n {
			rowMatch = true
			break
		}
	}

	if rowMatch {
		return false, nil
	}

	// Check Group Validity
	var grpMatch bool
	rr := row - row%3
	cc := column - column%3
	for r := rr; r <= rr+2; r++ {
		for c := cc; c <= cc+2; c++ {
			if c == column && r == row {
				continue
			}
			if s[r][c] == n {
				grpMatch = true
				break
			}
		}
	}

	if grpMatch {
		return false, nil
	}

	return true, nil
}

// IsComplete - returns a boolean stating if this sudoku has blank spaces or not
func (s *SudokuT) IsComplete() bool {
	for r := range s {
		for c := range s[r] {
			if s[r][c] == 0 {
				return false
			}
		}
	}
	return true
}

// Solve - Solve the Sudoku using backtracking. Returns only a single solution
// A valid sudoku will be a 9x9 grid with numbers from 0 - 9 where the number 0 indicates a blank slot. Any other input will result in an error
// A good Sudoku has only one solution, although a sudoku can potentially have many solutions (For now we return the first found solution)
func Solve(s *SudokuT) bool {

	if s.IsComplete() {
		return true
	}

	for r := range s {
		for c := range s[r] {
			if s[r][c] == 0 {
				for num := 1; num <= 9; num++ {
					valid, err := s.ValidMove(r, c, num)
					if err != nil {
						panic(err)
					}
					if valid {
						s.SetElement(r, c, num)
						if Solve(s) {
							return true
						}
						s[r][c] = 0
					}
				}
				return false
			}
		}
	}
	return false
}
