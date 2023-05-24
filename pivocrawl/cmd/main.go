package main

import (
	"log"

	"github.com/melphi/pivocrawl/infra/deps"
)

func main() {
	log.Print("Loading dependecies")
	deps := deps.Init()

	log.Print("Starting API")
	deps.ApiService.Start()
}
