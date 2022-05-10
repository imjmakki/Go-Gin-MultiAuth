package main

import (
	"Multi-Auth/api"
	"Multi-Auth/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", api.Register)

	r.Run(":9090")

}
