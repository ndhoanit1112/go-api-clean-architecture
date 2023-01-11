package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/domain/entities"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/domain/services"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/infrastructure/repository"
)

type productHandlers struct{}

func (h *productHandlers) CreateProduct(repo *repository.ProductRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product entities.Product

		if err := ctx.ShouldBind(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Init ProductService
		service := services.InitProductService(repo)
		if err := service.CreateProduct(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"id": product.ID})
	}
}

func (h *productHandlers) GetProduct(repo *repository.ProductRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, id)
	}
}

func (h *productHandlers) GetProducts(repo *repository.ProductRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Init ProductService
		service := services.InitProductService(repo)
		products, err := service.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}
