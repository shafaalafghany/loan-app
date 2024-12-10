package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shafaalafghany/loan-app/model"
	"github.com/shafaalafghany/loan-app/repository"
	"github.com/shafaalafghany/loan-app/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Register(*fiber.Ctx, *model.UserRequest) error
	Login(*fiber.Ctx, *model.UserRequest) error
	Profile(*fiber.Ctx) error
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if exists != nil && exists.Email == body.Email {
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

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
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

func (u *UserService) Login(c *fiber.Ctx, data *model.UserRequest) error {
	u.log.Info("incoming request to login", zap.Any("data", data))

	exists, err := u.repo.GetByEmail(data.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Println("error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	if exists == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "email or password is wrong",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(exists.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "email or password is wrong",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  exists.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "email or password is wrong",
		})
	}

	response := model.UserLoginResponse{
		User:    *exists,
		Token:   token,
		Message: "login successfully",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *UserService) Profile(c *fiber.Ctx) error {
	u.log.Info("incoming request to get profile")
	userId, err := util.GetUserIdFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	data, err := u.repo.GetById(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)

}
