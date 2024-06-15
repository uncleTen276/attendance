package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uncleTen276/attendance.git/server/internal/configs"
	"github.com/uncleTen276/attendance.git/server/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// generatePasswordHash() use to generate hashed password
func GeneratePasswordHash(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// IsValidPassword() use to compare hash password
func IsValidPassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

// generateToken() use to create new token
func GenerateToken(user *models.User, during time.Duration) (string, error) {
	config := configs.AppConfig().Auth
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(during).Unix(),
	})

	return token.SignedString([]byte(config.JWTSecret))
}
