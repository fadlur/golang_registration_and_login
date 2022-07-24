package handler

import (
	"net/http"
	"registration/request"
	"registration/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userServiceInterface services.UserServiceInterface
}

func Inituserhandler(userServiceInterface services.UserServiceInterface) *userHandler {
	return &userHandler{userServiceInterface}
}

func (userHandler *userHandler) Register(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := userRequest.HashPassword(userRequest.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message":err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message":err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}





	user, err := userHandler.userServiceInterface.Create(userRequest)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":false,
			"message":err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":true,
		"message":nil,
		"data":user,
		"errors":nil,
	})
}

func (userHandler *userHandler) GetUserById(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message":err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	user, err := userHandler.userServiceInterface.FindByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message":err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":true,
		"message":nil,
		"data":user,
		"errors":nil,
	})
}

func (userHandler *userHandler) Login(ctx *gin.Context)  {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	user, err := userHandler.userServiceInterface.Login(userRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	if err := user.CheckPassword(userRequest.Password); err != nil {
		ctx.JSON(401, gin.H{
			"status": false,
			"message": "Unauthorized",
			"data":nil,
			"errors":nil,
		})
		return
	}
	
	token, err := services.GenerateJWT(int(user.ID), user.Email)
	if err != nil {
		ctx.JSON(401, gin.H{
			"status": false,
			"message": "Unauthorized",
			"data":nil,
			"errors":nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Login berhasil",
		"data": map[string]string {
			"access_token":token,
			"token_type":"Bearer",
		},
		"errors":nil,
	})
}