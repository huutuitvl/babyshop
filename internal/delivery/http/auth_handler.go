package http

import (
	"babyshop/internal/config"
	"babyshop/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(r *gin.Engine, db *gorm.DB) {
	h := &AuthHandler{DB: db}
	r.POST("/login", h.Login)
}

// Login handles user authentication by validating the provided email and password.
// It expects a JSON payload with "email" and "password" fields.
// If the payload is invalid, the user is not found, the account is not verified,
// or the password is incorrect, it responds with an appropriate error message.
// On successful authentication, it generates and returns a JWT token.
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var user domain.User
	if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if user.Status != 2 || user.VerifiedAt == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User has not been approved by admin",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := config.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
