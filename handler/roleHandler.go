package handler

import (
	"net/http"
	"registration/request"
	"registration/response"
	"registration/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	roleServiceInterface services.RoleServiceInterface
}

func InitRoleHandler(roleServiceInterface services.RoleServiceInterface) *roleHandler {
	return &roleHandler{roleServiceInterface}
}

func (roleHandler *roleHandler) GetRoles(ctx *gin.Context) {
	roles, err := roleHandler.roleServiceInterface.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"status":false,
			"message": err,
			"data":nil,
			"errors":nil,
		})

		return
	}

	var roleResponse []response.RoleResponse
	for _, r := range roles {
		res := response.ConvertRoleResponse(r)
		roleResponse = append(roleResponse, res)
	}
	ctx.JSON(http.StatusOK, gin.H {
		"status":true,
		"message":"Role found",
		"data":roleResponse,
		"errors":nil,
	})
}

func (roleHandler *roleHandler) CreateRole(ctx *gin.Context) {
	var roleRequest request.RoleRequest

	if err := ctx.ShouldBindJSON(&roleRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}
	role, err := roleHandler.roleServiceInterface.Create(roleRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	roleResponse := response.ConvertRoleResponse(role)
	ctx.JSON(http.StatusCreated, gin.H{
		"status":true,
		"message": "Role saved",
		"data":roleResponse,
		"errors":nil,
	})
}

func (roleHandler *roleHandler) GetRoleById(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	role, err := roleHandler.roleServiceInterface.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}
	
	if role.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":false,
			"message": "Role not found",
			"data": nil,
			"errors":nil,
		})
		return
	}

	

	res := response.ConvertRoleResponse(role)
	ctx.JSON(http.StatusOK, gin.H{
		"status":true,
		"message": "Role found",
		"data": res,
		"errors":nil,
	})
}

func (roleHandler *roleHandler) DeleteRole(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	newRole, err := roleHandler.roleServiceInterface.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":false,
			"message":err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	roleResponse := response.ConvertRoleResponse(newRole)
	ctx.JSON(http.StatusOK, gin.H{
		"status":true,
		"message":"Role deleted",
		"data":roleResponse,
		"errors":nil,
	})
}

func (roleHandler *roleHandler) UpdateRole(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}
	var roleRequest request.RoleRequest

	if err := ctx.ShouldBindJSON(&roleRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}
	newRole, err := roleHandler.roleServiceInterface.Update(roleRequest, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"message": err.Error(),
			"data":nil,
			"errors":nil,
		})
		return
	}

	roleResponse := response.ConvertRoleResponse(newRole)
	ctx.JSON(http.StatusOK, gin.H{
		"status":false,
		"message":"Role updated",
		"data":roleResponse,
		"errors":nil,
	})
}