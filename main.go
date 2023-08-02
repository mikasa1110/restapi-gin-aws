package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mikasa1110/restapi-gin-gorm-docker/config"
	"github.com/mikasa1110/restapi-gin-gorm-docker/router"
)

func main() {
	r := gin.Default()
	config.Connect()
	router.UserRoute(r)
	r.Run(":8080")
}
