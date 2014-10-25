package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", rootController)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8080", nil)
}

func rootController(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/index.html")
}
