package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/example/fullstack-template/internal/config"
	"github.com/example/fullstack-template/internal/db"
	"github.com/example/fullstack-template/internal/models"
	"github.com/example/fullstack-template/internal/repository"
	"github.com/example/fullstack-template/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	action := os.Args[1]
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	gormDB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	if gormDB == nil {
		log.Fatal("db: connection is nil")
	}

	switch action {
	case "migrate":
		if err := gormDB.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migrate: %v", err)
		}
		fmt.Println("migrate: ok")
	case "seed":
		if err := gormDB.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migrate: %v", err)
		}
		authSvc := service.NewAuthService(repository.NewUserRepository(gormDB))
		if err := authSvc.SeedInitialUser(context.Background(), cfg.SeedEmail, cfg.SeedPassword); err != nil {
			log.Fatalf("seed: %v", err)
		}
		fmt.Printf("seed: ok (email=%s)\n", cfg.SeedEmail)
	case "bootstrap":
		if err := gormDB.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migrate: %v", err)
		}
		authSvc := service.NewAuthService(repository.NewUserRepository(gormDB))
		if err := authSvc.SeedInitialUser(context.Background(), cfg.SeedEmail, cfg.SeedPassword); err != nil {
			log.Fatalf("seed: %v", err)
		}
		fmt.Printf("bootstrap: migrate + seed complete (email=%s)\n", cfg.SeedEmail)
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("usage: go run ./cmd/dbtool [migrate|seed|bootstrap]")
}
