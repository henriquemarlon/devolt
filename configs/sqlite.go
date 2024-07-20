package configs

import (
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupSQlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	err = db.AutoMigrate(
		&entity.Bid{},
		&entity.User{},
		&entity.Order{},
		&entity.Auction{},
		&entity.Station{},
		&entity.Contract{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, nil
}
