package main

import (
	"log"
	"net/http"

	server "github.com/tatoalonso/apiquizgo/pkg/infrastructure"
)

func main() {

	r := server.NewServer()

	log.Fatal(http.ListenAndServe(":8000", r.Router))
}
