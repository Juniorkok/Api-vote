package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func InitUser(ctx *gin.Context) {
	r := gin.Default()
	r.GET("/users/:uuid", GetUserHandler)
	r.DELETE("/users/:uuid", DeleteUserHandler)
	r.PUT("/users/:uuid", PutUserHandler)
	r.GET("/users", GetAllUserHandler)
	r.POST("/users", PostUserHandler)
}

func GetUserHandler(ctx *gin.Context) {
	if u, ok := db.ListUser[ctx.Param("uuid")]; ok {
		ctx.JSON(http.StatusOK, u)
		return
	}
	ctx.JSON(http.StatusNotFound, nil)
}

func GetAllUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, listUser)
}

func PostUserHandler(ctx *gin.Context) {
	var u User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.UUID = uuid.NewV4().String()
	listUser[u.UUID] = u
	ctx.JSON(http.StatusOK, u)
}

func DeleteUserHandler(ctx *gin.Context) {
	if u, ok := db.ListUser[ctx.Param("uuid")]; ok {
		delete(listUser, ctx.Param("uuid"))
		ctx.JSON(http.StatusOK, u)
		return
	}

	ctx.JSON(http.StatusNotFound, nil)
}

func PutUserHandler(ctx *gin.Context) {
	u, ok := listUser[ctx.Param("uuid")]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	var newUser User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	u.FirstName = newUser.FirstName
	u.LastName = newUser.LastName
	db.ListUser[ctx.Param("uuid")] = u
	ctx.JSON(http.StatusOK, u)
}