package services

import (
	"errors"

	"github.com/ndhoanit1112/go-api-clean-architecture/internal/domain/entities"
	"gorm.io/datatypes"
)

// Interface provided by Business layer
// and will be implemented by Storage/Infrastructure layer
type ProductRepository interface {
	CreateProduct(data *entities.Product) error
	GetProducts() ([]entities.Product, error)
}

// Service provided by Business layer
// and will be used in Handlers (Presentation layer)
// NOTE: Private declaration for encapsulation
type productService struct {
	repository ProductRepository
}

// Public
// Provide access to other packages
func InitProductService(repo ProductRepository) *productService {
	return &productService{repository: repo}
}

// Define a business logic to service
func (service *productService) CreateProduct(data *entities.Product) error {
	// The business rules are implemented here
	// This is the core and the most important part of the system

	// Validation ...
	if data.Name == "" {
		return errors.New("Product name cannot be empty")
	}

	// Process data following business rules...
	data.Config = datatypes.JSON([]byte(`{"deploy_version": "v4"}`))

	// Delegate the work of Inserting product record into database to the storage layer
	// This kind of work is called "detail" in the context of Clean Architecture
	// Our business logic doesn't need to know anything about "details"
	// Note that we are communicate with storage layer via interface
	// and "details" are the implementation of interfaces
	if err := service.repository.CreateProduct(data); err != nil {
		return err
	}

	return nil
}

// Define a business logic to service
func (service *productService) GetProducts() ([]entities.Product, error) {
	products, err := service.repository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
