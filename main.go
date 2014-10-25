package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", rootController)
	http.ListenAndServe(":8080", nil)
}

func rootController(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main/index.html")
}
