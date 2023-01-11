package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/presentation/handlers/product"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	product.SetupRoutes(r, db)
}
