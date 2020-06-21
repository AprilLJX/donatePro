package main

import (
	"Donate_gin/db"
)

func main() {
	defer db.DB.Close()
	router := initRouter()
	router.Run(":9090")
}