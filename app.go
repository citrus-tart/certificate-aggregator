package main

import (
	"log"

	"github.com/citrus-tart/certificate-aggregator/api"
	"github.com/citrus-tart/certificate-aggregator/events"
	"github.com/citrus-tart/certificate-aggregator/processor"
	"github.com/citrus-tart/certificate-aggregator/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := repository.New()
	p := processor.New(r)

	events.CatchUp(p, "")

	api.Init(r)
}
