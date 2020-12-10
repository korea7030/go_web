package myapp

import (
	"net/http"
	"strconv"
	"web1/model"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"
)

var rd = render.New()

type Apphandler struct {
	http.Handler
	db model.DBHandler
}

func (a *Apphandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *Apphandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *Apphandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *Apphandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)

	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *Apphandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	completed := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, completed)

	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *Apphandler) Close() {
	a.db.Close()
}

func MakeHandler(filepath string) *Apphandler {
	r := mux.NewRouter()
	a := &Apphandler{
		Handler: r,
		db:      model.NewDBHandler(filepath),
	}
	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	return a
}
