package entities

import "gorm.io/datatypes"

type Product struct {
	BaseEntity
	Name                string
	Config              datatypes.JSON
	MaxFailureCountUser *int
	MaxFailureTimeUser  *int
}

func (Product) TableName() string {
	return "products"
}
