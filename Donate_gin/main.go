package main

import (
	"Donate_gin/db"
	"Donate_gin/models"
)

func main() {
	models.Code = make(map[string]string)
	defer db.DB.Close()
	router := initRouter()
	router.Run(":9090")
}