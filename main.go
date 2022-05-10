package main

import (
	"Multi-Auth/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", api.Register)

	r.Run(":9090")

}
