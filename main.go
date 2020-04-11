package main

import (
	"net/http"

	"github.com/nagarjun226/simplesudoku/api"
)

func main() {
	api := api.GetAPIInstance()
	r := api.Router()
	http.ListenAndServe(":4000", r)
}
