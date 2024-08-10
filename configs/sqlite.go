package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupSQlite() (*gorm.DB, error) {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(sqlite.Open("devolt.db"), &gorm.Config{
		Logger: logger,
	})
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

	db.Create(&entity.User{
		Role:      "admin",
		Address:   custom_type.NewAddress(common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")),
		CreatedAt: 0,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, nil
}

func SetupSQliteMemory() (*gorm.DB, error) {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger,
	})
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
	db.Create(&entity.User{
		Role:      "admin",
		Address:   custom_type.NewAddress(common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")),
		CreatedAt: 0,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, nil
}
