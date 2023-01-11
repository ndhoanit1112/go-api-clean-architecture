package entities

import "time"

type BaseEntity struct {
	ID        int
	CreatedAt *time.Time `gorm:"column:created"`
	UpdatedAt *time.Time `gorm:"column:updated"`
}
