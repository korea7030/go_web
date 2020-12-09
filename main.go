package main

import (
	"net/http"
	"web1/myapp"
)

func main() {
	http.ListenAndServe(":4000", myapp.NewHandler())
}
