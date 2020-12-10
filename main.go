package main

import (
	"log"
	"net/http"
	"web1/myapp"
)

func main() {
	m := myapp.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":4000", m)
	if err != nil {
		panic(err)
	}

}
