package api

import (
	"github.com/gorilla/mux"
)

// API - API instance
type API struct {
}

// GetAPIInstance - Return an Instance of the API
func GetAPIInstance() *API {
	a := new(API)
	return a
}

// Router - Router for the API
func (api *API) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/solve", sudokuSolverHandler).Methods("GET")
	return r
}
