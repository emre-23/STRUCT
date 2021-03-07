package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Ben Yunus Emre, n'aber?\n")
	fmt.Fprint(w, "<h1>Fenerbahçe</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
