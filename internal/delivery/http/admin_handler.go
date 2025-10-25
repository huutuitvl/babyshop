package http

import (
	"babyshop/internal/domain"
	"babyshop/internal/infrastructure"
	// "babyshop/internal/module/invoice/pdf"
	"net/http"

	"babyshop/internal/utils"

	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
)

type AdminHandler struct {
	DB *gorm.DB
}

func NewAdminHandler(r *gin.Engine, db *gorm.DB) {
	h := &AdminHandler{DB: db}
	group := r.Group("/admin", JWTMiddleware(), RoleMiddleware("admin", "super_admin"))

	// User management
	group.GET("/users", h.ListUsers)
	group.PUT("/users/:id/verify", h.VerifyUser)

	// Product management
	group.GET("/products", h.ListProducts)
	group.POST("/products", h.CreateProduct)
	group.PUT("/products/:id", h.UpdateProduct)
	group.DELETE("/products/:id", h.DeleteProduct)

	// Order management
	group.GET("/orders", h.ListOrders)
	group.GET("/orders/:id", h.GetOrder)
	group.PUT("/orders/:id/status", h.UpdateOrderStatus)
	// group.GET("/orders/:id/export", h.ExportOrderPDF)


	superAdminGroup := r.Group("/superadmin", JWTMiddleware(), RoleMiddleware("super_admin"))
	superAdminGroup.PUT("/users/:id/role", h.UpdateUserRole)

	// Category management
	group.GET("/categories", h.ListCategories)
	group.POST("/categories", h.CreateCategory)
	group.PUT("/categories/:id", h.UpdateCategory)
	group.DELETE("/categories/:id", h.DeleteCategory)


}

func (h *AdminHandler) ListProducts(c *gin.Context) {
	var products []domain.Product
	h.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// CreateProduct handles the creation of a new product.
// It expects a JSON payload representing a product in the request body.
// If the slug is missing, it will be generated from the product name.
// The handler checks for slug uniqueness and validates minimum SEO requirements for meta fields.
// On success, the product is saved to the database and returned in the response.
// Returns HTTP 400 for invalid payload, duplicate slug, or insufficient SEO meta fields.
// Returns HTTP 500 if database creation fails.
func (h *AdminHandler) CreateProduct(c *gin.Context) {
	var req domain.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Generate slug if not provided
	if req.Slug == "" {
		req.Slug = utils.GenerateSlug(req.Name)
	}

	// Check if slug already exists
	var exists int64
	h.DB.Model(&domain.Product{}).Where("slug = ?", req.Slug).Count(&exists)
	if exists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slug already exists"})
		return
	}

	// Validate minimum SEO requirements
	if len(req.MetaTitle) < 10 || len(req.MetaDescription) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Meta fields too short for SEO"})
		return
	}

	adminID := c.GetUint("user_id")
	req.CreatedByID = &adminID
	req.UpdatedByID = &adminID

	if err := h.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// UpdateProduct handles updating an existing product by its ID.
func (h *AdminHandler) UpdateProduct(c *gin.Context) {
	var existing domain.Product
	if err := h.DB.First(&existing, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	var req domain.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	adminID := c.GetUint("user_id")

	existing.Name = req.Name
	existing.Price = req.Price
	existing.Size = req.Size
	existing.Stock = req.Stock
	existing.ImageURL = req.ImageURL
	existing.UpdatedByID = &adminID

	h.DB.Save(&existing)
	c.JSON(http.StatusOK, existing)
}

func (h *AdminHandler) DeleteProduct(c *gin.Context) {
	var product domain.Product
	if err := h.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	adminID := c.GetUint("user_id")
	product.DeletedByID = &adminID
	h.DB.Save(&product)

	h.DB.Delete(&product) // Soft delete
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func (h *AdminHandler) ListOrders(c *gin.Context) {
	var orders []domain.Order
	h.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func (h *AdminHandler) GetOrder(c *gin.Context) {
	var order domain.Order
	if err := h.DB.Preload("Items.Product").First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *AdminHandler) UpdateOrderStatus(c *gin.Context) {
	var req struct {
		Status string `json:"status"` // ex: "paid", "cancelled", "shipped"
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var order domain.Order
	if err := h.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	adminID := c.GetUint("user_id")
	order.Status = req.Status
	order.UpdatedByID = &adminID
	h.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}

func (h *AdminHandler) UpdateUserRole(c *gin.Context) {
	var user domain.User
	id := c.Param("id")
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var req struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	user.Role = req.Role
	h.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Role updated"})
}

func (h *AdminHandler) ListCategories(c *gin.Context) {
	var cats []domain.Category
	h.DB.Find(&cats)
	c.JSON(http.StatusOK, cats)
}

func (h *AdminHandler) CreateCategory(c *gin.Context) {
	var req domain.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	req.Slug = utils.GenerateSlug(req.Name)
	adminID := c.GetUint("user_id")
	req.CreatedByID = &adminID
	req.UpdatedByID = &adminID

	h.DB.Create(&req)
	c.JSON(http.StatusCreated, req)
}

func (h *AdminHandler) UpdateCategory(c *gin.Context) {
	var cat domain.Category
	if err := h.DB.First(&cat, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var req domain.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	cat.Name = req.Name
	cat.Description = req.Description
	cat.Slug = utils.GenerateSlug(req.Name)
	adminID := c.GetUint("user_id")
	cat.UpdatedByID = &adminID

	h.DB.Save(&cat)
	c.JSON(http.StatusOK, cat)
}

func (h *AdminHandler) DeleteCategory(c *gin.Context) {
	var cat domain.Category
	if err := h.DB.First(&cat, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	h.DB.Delete(&cat)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func (h *AdminHandler) UploadProductImage(c *gin.Context) {
	productID := c.Param("id")

	// Lấy file
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "Missing file"})
		return
	}

	uploader, err := infrastructure.NewS3Uploader()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect S3"})
		return
	}

	url, err := uploader.UploadFile(file, fileHeader, "products/"+productID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Upload failed"})
		return
	}

	// Save DB
	uintID, err := utils.StringToUint(productID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	img := domain.ProductImage{
		ProductID: uintID,
		URL:       url,
		IsDefault: false,
	}

	adminID := c.GetUint("user_id")
	img.CreatedByID = &adminID
	img.UpdatedByID = &adminID

	h.DB.Create(&img)
	c.JSON(200, img)
}

// func (h *AdminHandler) ExportOrderPDF(c *gin.Context) {
// 	orderID := c.Param("id")
// 	var order domain.Order

// 	if err := h.DB.Preload("User").Preload("Items.Product").First(&order, orderID).Error; err != nil {
// 		c.JSON(404, gin.H{"error": "Order not found"})
// 		return
// 	}

// 	data := pdf.InvoiceViewData{
// 		OrderCode:    fmt.Sprintf("INV%05d", order.ID),
// 		CustomerName: order.User.Name,
// 		Address:      order.User.Address,
// 		Phone:        order.User.Phone,
// 		TotalAmount:  order.TotalPrice,
// 	}

// 	for _, item := range order.Items {
// 		data.Items = append(data.Items, pdf.InvoiceItem{
// 			Name:     item.Product.Name,
// 			Price:    item.Price,
// 			Quantity: item.Quantity,
// 			Subtotal: item.Subtotal,
// 		})
// 	}

// 	pdfPath, err := pdf.ExportInvoice(data)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.FileAttachment(pdfPath, filepath.Base(pdfPath))
// }



