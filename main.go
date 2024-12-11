package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shafaalafghany/loan-app/handler"
	"github.com/shafaalafghany/loan-app/middleware"
	"github.com/shafaalafghany/loan-app/model"
	"github.com/shafaalafghany/loan-app/repository"
	"github.com/shafaalafghany/loan-app/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBName     string
	JwtSecret  string
	AppPort    string
}

func main() {
	_ = godotenv.Load()

	config := Config{
		JwtSecret:  os.Getenv("SECRET_KEY"),
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
	}

	logConfig := zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(&model.User{}, &model.Transaction{}, &model.Limit{}, &model.AuditLog{})

	auditLogRepo := repository.NewAuditLogRepository(db, logger)

	limitRepo := repository.NewLimitRepository(db, logger)
	limitService := service.NewLimitService(limitRepo, logger)

	userRepo := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepo, auditLogRepo, logger)
	userHandler := handler.NewUserHandler(userService, logger)

	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	midAuth := middleware.JWTMiddleware(config.JwtSecret)

	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	api.Get("/profile", midAuth, userHandler.Profile)

	port := fmt.Sprintf(":%s", config.AppPort)
	log.Println("Server is running on port ", config.AppPort)
	app.Listen(port)
}
