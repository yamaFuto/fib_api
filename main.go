package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
}

func main() {
	var app application

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
