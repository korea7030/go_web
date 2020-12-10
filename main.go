package main

import (
	// "html/template"
	"os"
	"text/template"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "aaa", Email: "bbb", Age: 23}
	user2 := User{Name: "bbb", Email: "aaa@naver.com", Age: 40}
	users := []User{user, user2}
	// template 파일
	templ, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	// template 파일로 하지 않을 경우
	// templ.Execute(os.Stdout, user)
	// templ.Execute(os.Stdout, user2)

	// template파일을 사용할 경우
	templ.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
