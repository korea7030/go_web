package main

import (
	"net/http"

	"./myapp"
)

func main() {
	http.ListenAndServe(":4000", myapp.NewHttpHandler)
}
