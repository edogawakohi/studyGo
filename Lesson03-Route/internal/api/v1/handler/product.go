package v1handler

import (
	"main/utils"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

var (
	slugRegex   = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductV1(ctx *gin.Context) {

	search := ctx.Query("search")

	if err := utils.ValidationRequired("search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.ValidationStringLength("search", search, 3, 50); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errorMess := "Search must contain only letters, numbers and spaces"
	if err := utils.ValidationRegexSearch("search", search, searchRegex, errorMess); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limitStr)

	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Limit must be a positive numberSS",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products(ver1)",
		"search":  search,
		"limit":   limit})
}

func (u *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {

	slug := ctx.Param("slug")
	errorMess := "must contain only lowupper case letter, number, hyphens and dots"
	if err := utils.ValidationRegexSearch("slug", slug, slugRegex, errorMess); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Slug(ver1)",
		"slug":    slug,
	})
}

func (u *ProductHandler) PostProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create Product(ver1)"})
}
func (u *ProductHandler) PutProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update product (ver1)"})
}
func (u *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Update product (ver1)"})
}
