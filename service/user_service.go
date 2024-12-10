package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shafaalafghany/loan-app/model"
	"github.com/shafaalafghany/loan-app/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(*fiber.Ctx, *model.UserRequest) error
}

type UserService struct {
	repo repository.UserRepositoryInterface
	log  *zap.Logger
}

func NewUserService(repo repository.UserRepositoryInterface, log *zap.Logger) UserServiceInterface {
	return &UserService{
		repo: repo,
		log:  log,
	}
}

func (u *UserService) Register(c *fiber.Ctx, body *model.UserRequest) error {
	u.log.Info("incoming request to register new user", zap.Any("data", body))

	exists, err := u.repo.GetByEmail(body.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if exists.Email == body.Email {
		u.log.Error("email already exists")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	data := &model.User{
		ID:           uuid.New(),
		Email:        body.Email,
		NIK:          body.NIK,
		FullName:     body.FullName,
		LegalName:    body.LegalName,
		TempatLahir:  body.TempatLahir,
		TanggalLahir: body.TanggalLahir,
		Gaji:         body.Gaji,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("failed to hash password", zap.Any("error", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to hash password",
		})
	}

	data.Password = string(hash)

	if err := u.repo.Create(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create new user data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "create new user successfully",
	})
}
