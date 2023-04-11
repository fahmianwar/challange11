package repository

import "challange11/models"

type ProductRepository interface {
	FindById(id int) *models.Product
	GetAll(found bool) []models.Product
}
