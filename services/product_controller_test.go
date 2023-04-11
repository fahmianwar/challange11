package services

import (
	"challange11/models"
	"challange11/repository"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", 0).Return(nil)

	product := productService.GetOneProduct(0)

	assert.Nil(t, product)
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := models.Product{
		Title: "Coba", Description: "Coba", UserID: 1, User: nil,
	}

	productRepository.Mock.On("FindById", 1).Return(&product)

	result := productService.GetOneProduct(1)

	assert.NotNil(t, result)
	assert.Equal(t, &product, result, "result has to be a product data with id '1'")
}

func TestProductServiceGetProductsNotFound(t *testing.T) {
	productRepository.Mock.On("GetAll")

	product := productService.GetAll(false)

	assert.Nil(t, product)
}

func TestProductServiceGetProducts(t *testing.T) {
	products := []models.Product{
		{Title: "Coba", Description: "Coba", UserID: 1, User: nil},
		{Title: "Coba1", Description: "Coba1", UserID: 2, User: nil},
		{Title: "Coba2", Description: "Coba2", UserID: 3, User: nil},
	}

	productRepository.Mock.On("GetAll")

	result := productService.GetAll(true)

	assert.NotNil(t, result)
	assert.Equal(t, &products, &result, "result has to be a product data")
}
