package webserver

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/", home)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from /hello"))
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
