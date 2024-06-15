package middlewares

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uncleTen276/attendance.git/server/internal/configs"
	"github.com/uncleTen276/attendance.git/server/pkg/models"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() func(*fiber.Ctx) error {
	jwtwareConfig := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(configs.AppConfig().Auth.JWTSecret),
		},
		ContextKey:     "user",
		ErrorHandler:   jwtError,
		SuccessHandler: verifyTokenExpiration,
		TokenLookup:    "header:Authorization",
		AuthScheme:     "",
	}
	return jwtware.New(jwtwareConfig)
}

func verifyTokenExpiration(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	expires := int64(claims["exp"].(float64))
	if time.Now().Unix() > expires {
		return jwtError(c, errors.New("token expired"))
	}

	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
		Message: err.Error(),
	})
}

// GetUserIdFromToken() return userId from token
func GetUserIdFromToken(token *jwt.Token) (*int, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("can't extract user info from request")
	}
	var uid int = int(userId)

	return &uid, nil
}
