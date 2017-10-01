package main

import (
	"fmt"
	"net/http"
)

func placeHolder() {
    fmt.Println("Hello Davids World")
}

func main() {
	http.HandleFunc("/", placeHolder)
	http.ListenAndServe(":8080", nil)
}

