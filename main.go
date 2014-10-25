package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", rootController)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":8080", nil)
}

func rootController(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
