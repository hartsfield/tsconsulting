package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.tmpl", &page{Title: "Title", Body: "Hello world"})
	if err != nil {
		fmt.Println(err)
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	cd, _ := marshalContactData(r)
	fmt.Println(cd)
}
