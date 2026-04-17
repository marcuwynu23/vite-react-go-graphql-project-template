package db

import (
	"log"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	dsn := strings.TrimSpace(databaseURL)
	if dsn == "" {
		dsn = "app.db"
		log.Printf("db: DB_URL not set; defaulting to sqlite (%s)", dsn)
	}

	dialector := sqlite.Open(dsn)
	if strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://") {
		dialector = postgres.Open(dsn)
	}

	g, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := g.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	return g, nil
}

func DriverName(databaseURL string) string {
	dsn := strings.TrimSpace(databaseURL)
	if strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://") {
		return "postgres"
	}
	return "sqlite"
}
