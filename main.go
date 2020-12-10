package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/urfave/negroni"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "aaa", Email: "bbb"}

	rd.JSON(w, http.StatusOK, user)
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))

}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "aaa", Email: "bbb"}
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusInternalServerError)
	// 	// fmt.Fprint(w, err)
	// 	rd.Text(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	rd.HTML(w, http.StatusOK, "body", user)
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "aaa")
}

func main() {
	// templates확장자가 기본 폴더: "templates", 파일확장자: "tmpl" 로 되어있음
	rd = render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	// mux.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":4000", n)
}
