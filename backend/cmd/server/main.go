package main

import (
	"context"
	"log"
	"strings"

	"github.com/example/fullstack-template/internal/config"
	"github.com/example/fullstack-template/internal/db"
	graphql2 "github.com/example/fullstack-template/internal/graphql"
	"github.com/example/fullstack-template/internal/graphql/generated"
	"github.com/example/fullstack-template/internal/middleware"
	"github.com/example/fullstack-template/internal/models"
	"github.com/example/fullstack-template/internal/repository"
	"github.com/example/fullstack-template/internal/service"
	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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
	if err := gormDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	userRepo := repository.NewUserRepository(gormDB)
	authSvc := service.NewAuthService(userRepo)
	if err := authSvc.SeedInitialUser(context.Background(), cfg.SeedEmail, cfg.SeedPassword); err != nil {
		log.Fatalf("seed user: %v", err)
	}
	gqlSrv := gqlhandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql2.Resolver{Auth: authSvc},
	}))

	app := fiber.New(fiber.Config{
		AppName: "fullstack-template",
	})

	app.Use(recover.New())
	app.Use(middleware.RequestLogger())
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
		AllowOriginsFunc: allowOrigin(cfg.CORSOrigins),
	}))

	app.Get("/playground", adaptor.HTTPHandler(playground.Handler("GraphQL", "/graphql")))
	app.Get("/graphql", adaptor.HTTPHandler(playground.Handler("GraphQL", "/graphql")))
	app.Post("/graphql", adaptor.HTTPHandler(gqlSrv))

	log.Printf("listening on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func allowOrigin(origins []string) func(string) bool {
	allowed := make(map[string]struct{}, len(origins))
	for _, o := range origins {
		trimmed := strings.TrimSpace(o)
		if trimmed != "" {
			allowed[trimmed] = struct{}{}
		}
	}

	return func(origin string) bool {
		if _, ok := allowed[origin]; ok {
			return true
		}
		return strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:")
	}
}
