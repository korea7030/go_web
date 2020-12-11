package main

import (
	"log"
	"net/http"
	"os"
	"web/todos/myapp"
)

func main() {
	port := os.Getenv("PORT")
	m := myapp.MakeHandler(os.Getenv("DATABASE_URL"))
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":"+port, m)
	if err != nil {
		panic(err)
	}

}
