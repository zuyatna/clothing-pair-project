package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/repository"
)

type ProductDetailService struct {
	ProductDetailService repository.ProductDetailRepository
}

func NewProductDetailService(productDetailRepository repository.ProductDetailRepository) *ProductDetailService {
	return &ProductDetailService{ProductDetailService: productDetailRepository}
}

func (service *ProductDetailService) GetAllProductDetails() ([]models.ProductDetail, error) {
	return service.ProductDetailService.FindAll()
}

func (service *ProductDetailService) GetProductDetailByID(id int) (models.ProductDetail, error) {
	return service.ProductDetailService.FindByID(id)
}

func (service *ProductDetailService) AddProductDetail(productDetail models.ProductDetail) error {
	return service.ProductDetailService.Add(productDetail)
}

func (service *ProductDetailService) UpdateProductDetail(productDetail models.ProductDetail) error {
	return service.ProductDetailService.Update(productDetail)
}

func (service *ProductDetailService) DeleteProductDetail(id int) error {
	return service.ProductDetailService.Delete(id)
}
