package product

import (
	"github.com/gin-gonic/gin"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	h := &productHandlers{}
	repo := repository.InitProductRepo(db)

	routes := r.Group("/products")
	routes.GET("/", h.GetProducts(repo))
	routes.POST("/", h.CreateProduct(repo))
	routes.GET("/:id", h.GetProduct(repo))
}
