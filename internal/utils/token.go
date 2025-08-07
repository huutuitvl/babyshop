package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Load JWT secret from environment variable at runtime.
// It's better to load this once in an init() function for reliability.
var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Panic or handle error as needed in production
		panic("JWT_SECRET environment variable not set")
	}
	jwtSecret = []byte(secret)
}

// CustomClaims defines the structure of JWT claims used in the app.
type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token for a given user ID and role.
func GenerateToken(userID uint, role string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(userID)), // Optional: user ID as subject
			Issuer:    "babyshop-backend",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // Token expires in 72 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with claims and sign it using HS256 and the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken validates a JWT token string and returns the claims if valid.
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	// Check for parsing errors
	if err != nil {
		// Handle token expiration error specifically
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	// Validate token and extract claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
