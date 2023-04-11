package repository

import (
	"challange11/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id int) *models.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := models.Product{
		Title: "Coba", Description: "Coba", UserID: 1, User: nil,
	}
	return &product
}

func (repository *ProductRepositoryMock) GetAll(found bool) []models.Product {
	// db := database.GetDB()
	products := []models.Product{
		{Title: "Coba", Description: "Coba", UserID: 1, User: nil},
		{Title: "Coba1", Description: "Coba1", UserID: 2, User: nil},
		{Title: "Coba2", Description: "Coba2", UserID: 3, User: nil},
	}

	if !found {
		return nil
	} else {
		return products
	}

	// if err := db.Debug().Find(&products).Error; err != nil {
	// 	return nil
	// }

	// products[0].ID = 1
	// products[0].Title = "Coba"
	// products[0].Description = "Cobalah"
	// products[1].ID = 2
	// products[1].Title = "Coba1"
	// products[1].Description = "Cobalah1"

	// if len(products) > 1 {
	// 	return products
	// }
	// return nil

	// err := db.Find(products).Error

}
