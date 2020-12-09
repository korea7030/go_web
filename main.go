package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Bar!")
}

func main() {
	// 경로에 해당하는 handlerFunc 설정
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") // writer에 해당하는 문자열을 출력
	})

	http.HandleFunc("/bar", barHandler)

	// interface 호출
	http.Handle("/foo", &fooHandler{})

	// Server Listening
	http.ListenAndServe(":4000", nil)
}
