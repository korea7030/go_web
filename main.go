package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	// annotation 추가
	FirstName string    `mytag:"first_name"`
	LastName  string    `mytag:"last_name"`
	Email     string    `mytag:"email"`
	createdAt time.Time `mytag:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request", err)
		return
	}
	user.createdAt = time.Now()
	fmt.Print("time : ", user)

	// go structure -> json 형태로 변환(bytearray, err)
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	// string 형태로 변환
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // url argument

	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	mux := http.NewServeMux() // route instance 생성

	// 경로에 해당하는 handlerFunc 설정
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") // writer에 해당하는 문자열을 출력
	})

	mux.HandleFunc("/bar", barHandler)

	// interface 호출
	mux.Handle("/foo", &fooHandler{})

	// Server Listening
	http.ListenAndServe(":4000", mux)
}
