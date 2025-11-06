package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default() // tao server có sẵn middleware
	// có recovery và logger, còn một cách khác sẽ tạo ra mà không có recovery và logger gin.New()
	//create endpoint
	router.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Kohi pro Golang"})
	})

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Golang cung binh thuong"})
	})

	router.GET("/user/:user_id", func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")

		salary := ctx.Query("salary")

		color := ctx.Query("color")

		ctx.JSON(200, gin.H{
			"user_id": user_id,
			"data":    "Thong tin Golang",
			"salary":  salary,
			"color":   color})
	})
	router.Run(":8080") //start server
}
