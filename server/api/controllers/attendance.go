package controllers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	middlewares "github.com/uncleTen276/attendance.git/server/api/middleware"
	"github.com/uncleTen276/attendance.git/server/internal/contracts"
	"github.com/uncleTen276/attendance.git/server/internal/postgresql"
	"github.com/uncleTen276/attendance.git/server/pkg/models"
	"github.com/uncleTen276/attendance.git/server/pkg/repositories"
	"github.com/uncleTen276/attendance.git/server/pkg/validators"
)

// AttendanceRecordCheckIn()
// @AttendanceRecordCheckIn godoc
// @Summary api use for employee checkIn
// @Description check-in time should be format : 2024-06-15T08:00:00Z
// @Tags Attendance
// @Param todo body models.AttendanceCheckIn true "Check In position"
// @Accept json
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /attendance/ [POST]
// @Security Bearer
func AttendanceRecordCheckIn(c *fiber.Ctx) error {
	req := &models.AttendanceCheckIn{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err := validators.NewValidator().Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

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

	// check user existed
	userRepo := repositories.NewUserRepo(postgresql.GetDB())
	if _, err := userRepo.GetUserById(*userId); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	contract, client, err := contracts.GetContractInstance()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect contracts",
		})
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect contracts",
		})
	}

	auth, err := contracts.GetAuth(chainId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect Auth contracts",
		})
	}

	employeeId := big.NewInt(int64(*userId))
	tx, err := contract.RecordAtttendance(auth, employeeId, req.CheckIn, req.Detail, req.Position)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect Auth contracts",
		})
	}

	eventRepo := repositories.NewEventRepo(postgresql.GetDB())
	if err := eventRepo.Create(&models.CreateEvent{
		Hash:       tx.Hash().Hex(),
		Type:       repositories.CREATE_EVENT,
		EmployeeId: *userId,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happended",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GetAttendanceRecordByEmployee()
// @GetAttendanceRecordByEmployee godoc
// @Summary api use to get user attendance record
// @Tags Attendance
// @Accept json
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /attendance/  [GET]
// @Security Bearer
func GetAttendanceRecordByEmployee(c *fiber.Ctx) error {
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

	// check user existed
	userRepo := repositories.NewUserRepo(postgresql.GetDB())
	if _, err := userRepo.GetUserById(*userId); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	contract, _, err := contracts.GetContractInstance()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect contracts",
		})
	}

	attendances, err := contract.GetAttendanceByEmployee(nil, big.NewInt(int64(*userId)))
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect contracts",
		})
	}

	return c.JSON(models.BasePaginationResponse{
		Items:    attendances,
		Page:     1,
		PageSize: len(attendances),
	})
}

// UpdateAttendance()
// @UpdateAttendance godoc
// @Summary api use to update attendaceRecord
// @Description check-in time should be format : 2024-06-15T08:00:00Z
// @Tags Attendance
// @Param todo body models.AttendaceUpdate true "update record"
// @Accept json
// @Failure 400,403,404,500 {object} models.ErrorResponse "Error"
// @Router /attendance/ [PUT]
// @Security Bearer
func UpdateAttendance(c *fiber.Ctx) error {
	req := &models.AttendaceUpdate{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err := validators.NewValidator().Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Message: "invalid input found",
			Errors:  validators.ValidatorErrors(err),
		})
	}

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

	// check user existed
	userRepo := repositories.NewUserRepo(postgresql.GetDB())
	if _, err := userRepo.GetUserById(*userId); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(models.ErrorResponse{
			Message: "user not found",
		})
	}

	contract, client, err := contracts.GetContractInstance()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect contracts",
		})
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect contracts",
		})
	}

	auth, err := contracts.GetAuth(chainId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect Auth contracts",
		})
	}

	tx, err := contract.UpdateAttendance(auth, big.NewInt(int64(req.RecordId)), req.CheckIn, req.Detail, req.Position)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "can not connect Auth contracts",
		})
	}

	eventRepo := repositories.NewEventRepo(postgresql.GetDB())
	if err := eventRepo.Create(&models.CreateEvent{
		Hash:       tx.Hash().Hex(),
		Type:       repositories.UPDATE_EVENT,
		EmployeeId: *userId,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Message: "something bad happended",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
