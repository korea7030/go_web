package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	print(vars)
	fmt.Fprint(w, "User id ", vars["id"])
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo By /users/{id}")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// NewHandler
func NewHandler() http.Handler {
	mux := mux.NewRouter()
	// mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", UsersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", GetUserInfoHandler)
	return mux
}
