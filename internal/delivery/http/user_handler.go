package http

import (
	"babyshop/internal/usecase"
	"net/http"

	"babyshop/internal/domain"

	"github.com/gin-gonic/gin"

	"time"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, uc *usecase.UserUsecase) {
	h := &UserHandler{Usecase: uc}
	r.POST("/register", h.Register)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	err := h.Usecase.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func (h *AdminHandler) ListUsers(c *gin.Context) {
	var users []domain.User
	h.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

func (h *AdminHandler) VerifyUser(c *gin.Context) {
	var user domain.User
	id := c.Param("id")
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	now := time.Now()
	adminID := c.GetUint("user_id")

	user.Status = 2
	user.VerifiedAt = &now
	user.VerifiedByID = &adminID
	user.UpdatedByID = &adminID

	h.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User verified"})
}
