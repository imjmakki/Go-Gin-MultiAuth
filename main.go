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
	public.POST("/login", api.Login)

	r.Run(":9090")

}
