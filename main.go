package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hamhockwarlock/id-comparison/db"
	"github.com/hamhockwarlock/id-comparison/uuid"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := gin.Default()

	uuid.Routes(r)

	r.Run(":8080")
}
