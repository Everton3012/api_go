package controller

import "github.com/gin-gonic/gin"

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (pc *ProductController) GetProducts(c *gin.Context) {}
