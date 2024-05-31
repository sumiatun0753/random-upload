package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Uname, Pass string
}

func main() {

	login := User{
		Uname: "ain",
		Pass:  "ain",
	}

	tmp := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmp.Execute(w, nil)
			return
		}

		data := User{
			Uname: r.FormValue("uname"),
			Pass:  r.FormValue("pass"),
		}

		if login == data {
			tmp.Execute(w, struct{ Success bool }{true})
		} else {
			tmp.Execute(w, nil)
			return
		}

	})
	http.ListenAndServe(":8080", nil)

}
