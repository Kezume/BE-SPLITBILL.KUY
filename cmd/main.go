package main

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/migration"
	"github.com/gin-gonic/gin"
)

func main() {
	migration.ConnectDB()

	r := gin.Default()

	r.Run(":8080")
}
