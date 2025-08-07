package http

import (
	"babyshop/internal/domain"
	"babyshop/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Usecase *usecase.ProductUsecase
}

// NewProductHandler initializes the product handler with routes
// and binds the usecase to the handler methods.
// It sets up the HTTP methods for product operations.
func NewProductHandler(r *gin.Engine, uc *usecase.ProductUsecase) {
	handler := &ProductHandler{Usecase: uc}
	group := r.Group("/products", JWTMiddleware())

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.POST("", handler.Create)
	group.PUT("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
}

// GetAll handles GET requests to retrieve all products.
// It responds with a JSON array of products or an error message.
func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetByID handles GET requests to retrieve a product by its ID.
// It responds with the product in JSON format or an error message if not found.
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Create handles POST requests to create a new product.
// It binds the JSON payload to a Product struct and creates the product.
func (h *ProductHandler) Create(c *gin.Context) {
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Usecase.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

// Update handles PUT requests to update an existing product by its ID.
// It binds the JSON payload to a Product struct and updates the product.
func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.ID = uint(id)
	if err := h.Usecase.Update(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// Delete handles DELETE requests to remove a product by its ID.
// It deletes the product and responds with a success message or error.
func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
