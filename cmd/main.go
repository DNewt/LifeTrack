package main

import (
	"fmt"
	"net/http"
)

func placeHolder(w http.ResponseWriter, r *http.Request){
    fmt.Println("Hello Davids World")
}

func main() {
	http.HandleFunc("/", placeHolder)
	http.ListenAndServe(":8080", nil)
}

