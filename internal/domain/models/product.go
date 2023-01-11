package models

// All domain models related to Product entity

type ProductDetails struct {
	ID   string
	Name string
}

type ProductForCreate struct {
	ID                  string
	Name                string
	MaxFailureCountUser int
	MaxFailureTimeUser  int
}
