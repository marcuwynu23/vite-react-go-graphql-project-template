package main

import (
	"log"

	"github.com/example/fullstack-template/internal/config"
	"github.com/example/fullstack-template/internal/db"
	"github.com/example/fullstack-template/internal/handler"
	"github.com/example/fullstack-template/internal/middleware"
	"github.com/example/fullstack-template/internal/models"
	"github.com/example/fullstack-template/internal/repository"
	"github.com/example/fullstack-template/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	gormDB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	if gormDB != nil {
		if err := gormDB.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migrate: %v", err)
		}
	}

	userRepo := repository.NewUserRepository(gormDB)
	authSvc := service.NewAuthService(userRepo)
	authH := handler.NewAuthHandler(authSvc)

	app := fiber.New(fiber.Config{
		AppName: "fullstack-template",
	})

	app.Use(recover.New())
	app.Use(middleware.RequestLogger())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     joinOrigins(cfg.CORSOrigins),
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	app.Get("/health", handler.Health)
	app.Post("/api/auth/login", authH.Login)

	log.Printf("listening on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func joinOrigins(origins []string) string {
	if len(origins) == 0 {
		return "http://localhost:5173"
	}
	out := origins[0]
	for i := 1; i < len(origins); i++ {
		out += "," + origins[i]
	}
	return out
}
