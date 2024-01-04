package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World !")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
