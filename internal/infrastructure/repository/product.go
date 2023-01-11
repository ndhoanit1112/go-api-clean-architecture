package repository

import (
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func InitProductRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) CreateProduct(data *entities.Product) error {
	result := repo.db.Create(&data)
	return result.Error
}

func (repo *ProductRepository) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	result := repo.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
