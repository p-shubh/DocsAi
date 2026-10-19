package main

import (
	"log"

	"github.com/hexanecreations/docsai/backend/internal/bootstrap"
	"github.com/hexanecreations/docsai/backend/internal/server"
)

func main() {

	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.New(app)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.Run())
}
