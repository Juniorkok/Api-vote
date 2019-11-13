package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users/:uuid", GetUserHandler)
	r.DELETE("/users/:uuid", DeleteUserHandler)
	r.PUT("/users/:uuid", PutUserHandler)
	r.GET("/users", GetAllUserHandler)
	r.POST("/users", PostUserHandler)
	r.Run(":8080")
}
