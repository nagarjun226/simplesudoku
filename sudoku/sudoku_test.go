package sudoku

func returnTestCase(n int) interface{} {
	switch {
	case n == 1: // Complete Sudoku as a 1D array
		s := []int{4, 3, 5, 2, 6, 9, 7, 8, 1, 6, 8, 2, 5, 7, 1, 4, 9, 3, 1, 9, 7, 8, 3, 4, 5, 6, 2, 8, 2, 6, 1, 9, 5, 3, 4, 7, 3, 7, 4, 6, 8, 2, 9, 1, 5, 9, 5, 1, 7, 4, 3, 6, 2, 8, 5, 1, 9, 3, 2, 6, 8, 7, 4, 2, 4, 8, 9, 5, 7, 1, 3, 6, 7, 6, 3, 4, 1, 8, 2, 5, 9}
		return s
	case n == 2: // Incomplete Hard Sudoku with only one solution
		s := []int{1, 5, 8, 0, 6, 0, 0, 0, 0, 9, 0, 0, 0, 0, 4, 0, 7, 0, 3, 0, 0, 8, 0, 0, 0, 0, 0, 5, 0, 0, 2, 7, 0, 9, 0, 0, 0, 0, 0, 0, 3, 0, 0, 5, 7, 0, 7, 0, 0, 0, 0, 0, 1, 0, 4, 0, 0, 0, 0, 1, 8, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 1, 9, 0}
		return s
	}
	return nil
}
