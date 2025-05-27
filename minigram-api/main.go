package main

import (
	"minigram-api/repo"
	"minigram-api/routers"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))

	repo.StartDB()

	routers.StartServer().Run(":8080")
}
