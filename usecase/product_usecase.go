package usecase

import "api_go/model"

type ProductUsecase struct {
	//Repository
}

func NewProductUseCase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return nil, nil
}
