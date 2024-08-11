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
		Address:   custom_type.NewAddress(common.HexToAddress("0x0142f501EE21f4446009C3505c51d0043feC5c68")),
		CreatedAt: 0,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, nil
}