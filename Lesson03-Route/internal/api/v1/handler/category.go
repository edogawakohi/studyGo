package v1handler

import (
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

var validCategory = map[string]bool{
	"golang": true,
	"java":   true,
	"html":   true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryBySlugV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if err := utils.ValidationInList("category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}
