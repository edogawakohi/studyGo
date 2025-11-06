package main

import (
	v1handler "main/internal/api/v1/handler"
	v2package "main/internal/api/v2/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//router group
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerv1 := v1handler.NewUserHandler()
			user.GET("/", userHandlerv1.GetUserV1)
			user.GET("/:id", userHandlerv1.GetUserByIdV1)
			user.GET("admin/:uuid", userHandlerv1.GetUserByUUIdV1)
			user.POST("/", userHandlerv1.PostUserV1)
			user.PUT("/:id", userHandlerv1.PutUserV1)
			user.DELETE("/:id", userHandlerv1.DeleteUserV1)
		}

		product := v1.Group("/products")
		{
			productHandlerv1 := v1handler.NewProductHandler()

			product.GET("/", productHandlerv1.GetProductV1)
			product.GET("/:slug", productHandlerv1.GetProductBySlugV1)
			product.POST("/", productHandlerv1.PostProductV1)
			product.PUT("/:id", productHandlerv1.PutProductV1)
			product.DELETE("/:id", productHandlerv1.DeleteProductV1)
		}

		category := v1.Group("/categories")
		{
			categoryHandlerv1 := v1handler.NewCategoryHandler()
			category.GET("/:category", categoryHandlerv1.GetCategoryBySlugV1)
		}

	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			userHandlerv2 := v2package.NewUserHandler()

			user.GET(" ", userHandlerv2.GetUserV2)
			user.GET("/:id", userHandlerv2.GetUserByIdV2)
			user.POST(" ", userHandlerv2.PostUserV2)
			user.PUT("/:id", userHandlerv2.PutUserV2)
			user.DELETE("/id", userHandlerv2.DeleteUserV2)
		}

	}

	r.Run(":8080")

}
