package main

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	r.Run(":8080")
}
