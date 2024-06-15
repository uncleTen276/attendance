package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	middlewares "github.com/uncleTen276/attendance.git/server/api/middleware"
	"github.com/uncleTen276/attendance.git/server/internal/configs"
	"github.com/uncleTen276/attendance.git/server/internal/postgresql"
	"github.com/uncleTen276/attendance.git/server/pkg/models"
	"github.com/uncleTen276/attendance.git/server/pkg/repositories"
	"github.com/uncleTen276/attendance.git/server/pkg/util"
	"github.com/uncleTen276/attendance.git/server/pkg/validators"
)

// @CreateUser() godoc
// @Summary create new user
// @Tags Auth
// @Param todo body models.CreateUser true "New User"
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /register [post]
func CreateUser(c *fiber.Ctx) error {
	user := &models.CreateUser{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	userRepo := repositories.NewUserRepo(postgresql.GetDB())
	// find user by identifier
	userExist, err := userRepo.GetUserByIdentifier(user.Identifier)
	if userExist != nil {
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
			Message: "this user identifier already exists",
		})
	}

	if err != nil && err != models.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "fail to create user",
		})
	}

	hashedPassword, err := util.GeneratePasswordHash([]byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(1*time.Second))
	defer cancel()

	user.Password = hashedPassword
	err = userRepo.Create(user, ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
		})
	}

	return c.Status(http.StatusCreated).SendString("create success")
}

// Login
// @Login godoc
// @Summary User Login
// @Description use for login response the access_Token
// @Tags Auth
// @Accept json
// @Param todo body models.UserLogin true "Login"
// @Success 200 {object} models.UserToken
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	user := &models.UserLogin{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.ErrorResponse{
			Message: err.Error(),
		})
	}

	validate := validators.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

	userRepo := repositories.NewUserRepo(postgresql.GetDB())
	userExist, err := userRepo.GetUserByIdentifier(user.Identifier)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&models.ErrorResponse{
			Message: "user not found",
		})
	}

	if !util.IsValidPassword(userExist.Password, user.Password) {
		return c.Status(fiber.StatusForbidden).JSON(&models.ErrorResponse{
			Message: "wrong password",
		})
	}

	config := configs.AppConfig().Auth
	token, err := util.GenerateToken(userExist, time.Duration(config.TokenExpire*int(time.Minute)*24))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "some thing bad happended",
			Errors:  err.Error(),
		})
	}

	return c.JSON(models.UserToken{
		AccessToken: token,
	})
}

// GetUserInfo()
// @GetUserInfo godoc
// @Summary get user information by Id
// @Description route get user Id from token then get user information
// @Tags User
// @Accept json
// @Success 200 {object} models.User
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /user/me [get]
// @Security Bearer
func GetUserInfo(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(models.ErrorResponse{
			Message: "can not parse token",
		})
	}

	userId, err := middlewares.GetUserIdFromToken(token)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	userRepo := repositories.NewUserRepo(postgresql.GetDB())
	user, err := userRepo.GetUserById(*userId)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	return c.JSON(user)
}
