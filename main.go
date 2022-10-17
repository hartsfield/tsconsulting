package main

import (
	"context"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type page struct {
	Title string
	Body  string
}

type contactData struct {
	Email   string `json:"email"`
	Phone   string `json:"phoneNumber"`
	Message string `json:"message"`
}

var (
	templates = template.Must(template.New("main").ParseGlob("internal/*/*.tmpl"))
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/api/sendMail", sendMail)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Server configuration
	srv := &http.Server{
		// in production only ust SSL
		Addr:              ":9047",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	ctx, cancelCtx := context.WithCancel(context.Background())

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
		cancelCtx()
	}()

	fmt.Println("Server started @ " + srv.Addr)
	<-ctx.Done()
}
