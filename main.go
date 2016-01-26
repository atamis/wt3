package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", Routes())

}
