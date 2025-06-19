package main

import (
	"my_first_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterTaskRoutes(r)

	r.Run()
}
