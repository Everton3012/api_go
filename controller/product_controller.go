package controller

import (
	"api_go/model"
	"api_go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
