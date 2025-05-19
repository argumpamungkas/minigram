package main

import (
	"minigram-api/repo"
	"minigram-api/routers"
)

func main() {
	repo.StartDB()

	routers.StartServer().Run(":8080")
}
