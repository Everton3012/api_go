package main

import (
	"api_go/controller"
	"api_go/db"
	"api_go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	ProductUseCase := usecase.NewProductUseCase()

	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}
