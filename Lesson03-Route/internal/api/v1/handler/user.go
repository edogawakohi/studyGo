package v1handler

import (
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all users(ver1)"})
}

func (u *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := utils.ValidationPositiveInt("id", idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
		"user_id": id,
	})
}

func (u *UserHandler) GetUserByUUIdV1(ctx *gin.Context) {
	idStr := ctx.Param("uuid")

	uid, err := utils.ValidationUUId("uuid", idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
		"uuid":    uid,
	})
}

func (u *UserHandler) PostUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create User(ver1)"})
}
func (u *UserHandler) PutUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update user (ver1)"})
}
func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Update user (ver1)"})
}
