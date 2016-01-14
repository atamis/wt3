package main

import (
	"fmt"
	"net/http"
)

type State struct {
	On      bool
	Answers []uint
}

func main() {

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", Routes())

}
