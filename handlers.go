package main

import (
	"fmt"
	"net/http"

	"github.com/hartsfield/gmailer"
)

func home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "main.tmpl", &page{})
	if err != nil {
		fmt.Println(err)
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sendmail")
	m, _ := marshalContactData(r)
	msg := gmailer.Message{
		Recipient: "johnathanhartsfield@gmail.com",
		Subject:   "ALERT! NEW JOB REQUEST!",
		Body:      m.Email + "::" + m.Phone + "::" + m.Message,
	}
	msg.Send(onMessageSent)
	ajaxResponse(w, map[string]string{
		"success": "true",
	})
}

func onMessageSent() {
	fmt.Println("sent mail")
}
