package v2package

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all users(ver2)"})
}

func (u *UserHandler) GetUserByIdV2(ctx *gin.Context) {
	// id := ctx.Params("id")
	ctx.JSON(http.StatusOK, gin.H{"message": "Get User By Id(ver2)"})
}

func (u *UserHandler) PostUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create User(ver2)"})
}
func (u *UserHandler) PutUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update user (ver2)"})
}
func (u *UserHandler) DeleteUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Update user (ver2)"})
}
