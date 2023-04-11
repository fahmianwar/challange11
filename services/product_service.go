package services

import (
	"challange11/models"
	"challange11/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id int) *models.Product {
	product := service.Repository.FindById(id)
	return product
}

func (service ProductService) GetAll(found bool) []models.Product {
	product := service.Repository.GetAll(found)

	return product
}
