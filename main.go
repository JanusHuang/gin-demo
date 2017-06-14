package main

import (
	db "github.com/JanusHuang/gin-demo/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
