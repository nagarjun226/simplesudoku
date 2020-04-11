package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nagarjun226/simplesudoku/sudoku"
)

// SudokuAPIT - struct to handle the sudoku as a json
// provide API when input is a list or a list of lists
type SudokuAPIT struct {
	Sudoku []int `json:"sudoku"`
}

func sudokuSolverHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error readng request body")
		return
	}

	var su SudokuAPIT
	err = json.Unmarshal(reqBody, &su)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Processing Error")
		return
	}

	sud := su.Sudoku
	if len(sud) != 81 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Sudoku doesnt havr 81 elements (filled and blank combined) len = %v", len(sud))
		return
	}

	s := sudoku.GetSudoku()
	err = s.SetFromArray(sud)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Processing error %v", err)
		return
	}

	sudoku.Solve(s)

	ans, e := s.ArrayFormat()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Processing error %v", e)
		return
	}

	var resp SudokuAPIT
	resp.Sudoku = ans

	sJSON, er := json.Marshal(resp)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Processing error %v", er)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(sJSON))
}
