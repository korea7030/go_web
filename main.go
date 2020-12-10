package main

import (
	"log"
	"net/http"
	"web1/myapp"

	"github.com/urfave/negroni"
)

func main() {
	m := myapp.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":4000", n)
	if err != nil {
		panic(err)
	}

}
