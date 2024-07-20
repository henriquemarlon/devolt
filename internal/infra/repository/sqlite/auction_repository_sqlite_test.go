package sqlite

// import (
// 	"testing"
// 	"time"

// 	"github.com/devolthq/devolt/internal/domain/entity"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// 	"math/big"
// )

// var db *gorm.DB

// func TestMain(m *testing.M) {
// 	var err error
// 	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// Run migrations
// 	err = db.AutoMigrate(&entity.Auction{}, &entity.Bid{})
// 	if err != nil {
// 		panic("failed to run migrations")
// 	}

// 	m.Run()
// }

// func setupTestDB() error {
// 	// Drop and create tables for a clean state
// 	err := db.Migrator().DropTable(&entity.Auction{}, &entity.Bid{})
// 	if err != nil {
// 		return err
// 	}

// 	err = db.AutoMigrate(&entity.Auction{}, &entity.Bid{})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func TestAuctionRepositorySqlite(t *testing.T) {
// 	repo := NewAuctionRepositorySqlite(db)

// 	t.Run("CreateAuction", func(t *testing.T) {
// 		err := setupTestDB()
// 		assert.Nil(t, err)

// 		credits := big.NewInt(1000)
// 		priceLimit := big.NewInt(500)

// 		auction := &entity.Auction{
// 			Credits:    credits,
// 			PriceLimit: priceLimit,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		res, err := repo.CreateAuction(auction)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, res)
// 		assert.Equal(t, auction.Credits, res.Credits)
// 		assert.Equal(t, auction.PriceLimit, res.PriceLimit)
// 	})

// 	t.Run("FindActiveAuction", func(t *testing.T) {
// 		err := setupTestDB()
// 		assert.Nil(t, err)

// 		credits := big.NewInt(1000)
// 		priceLimit := big.NewInt(500)

// 		auction := &entity.Auction{
// 			Credits:    credits,
// 			PriceLimit: priceLimit,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		_, err = repo.CreateAuction(auction)
// 		assert.Nil(t, err)

// 		storedAuction, err := repo.FindActiveAuction()
// 		assert.Nil(t, err)
// 		assert.NotNil(t, storedAuction)
// 		assert.Equal(t, auction.Credits, storedAuction.Credits)
// 		assert.Equal(t, auction.PriceLimit, storedAuction.PriceLimit)
// 	})

// 	t.Run("FindAuctionById", func(t *testing.T) {
// 		err := setupTestDB()
// 		assert.Nil(t, err)

// 		credits := big.NewInt(1000)
// 		priceLimit := big.NewInt(500)

// 		auction := &entity.Auction{
// 			Credits:    credits,
// 			PriceLimit: priceLimit,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		res, err := repo.CreateAuction(auction)
// 		assert.Nil(t, err)

// 		storedAuction, err := repo.FindAuctionById(res.Id)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, storedAuction)
// 		assert.Equal(t, auction.Credits, storedAuction.Credits)
// 		assert.Equal(t, auction.PriceLimit, storedAuction.PriceLimit)
// 	})

// 	t.Run("FindAllAuctions", func(t *testing.T) {
// 		err := setupTestDB()
// 		assert.Nil(t, err)

// 		credits1 := big.NewInt(1000)
// 		priceLimit1 := big.NewInt(500)
// 		auction1 := &entity.Auction{
// 			Credits:    credits1,
// 			PriceLimit: priceLimit1,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		_, err = repo.CreateAuction(auction1)
// 		assert.Nil(t, err)

// 		credits2 := big.NewInt(2000)
// 		priceLimit2 := big.NewInt(1000)
// 		auction2 := &entity.Auction{
// 			Credits:    credits2,
// 			PriceLimit: priceLimit2,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		_, err = repo.CreateAuction(auction2)
// 		assert.Nil(t, err)

// 		storedAuctions, err := repo.FindAllAuctions()
// 		assert.Nil(t, err)
// 		assert.NotNil(t, storedAuctions)
// 		assert.Equal(t, 2, len(storedAuctions))
// 	})

// 	t.Run("UpdateAuction", func(t *testing.T) {
// 		err := setupTestDB()
// 		assert.Nil(t, err)

// 		credits := big.NewInt(1000)
// 		priceLimit := big.NewInt(500)

// 		auction := &entity.Auction{
// 			Credits:    credits,
// 			PriceLimit: priceLimit,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		_, err = repo.CreateAuction(auction)
// 		assert.Nil(t, err)

// 		auction.State = entity.AuctionFinished
// 		updatedAuction, err := repo.UpdateAuction(auction)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, updatedAuction)
// 		assert.Equal(t, entity.AuctionFinished, updatedAuction.State)
// 	})

// 	t.Run("DeleteAuction", func(t *testing.T) {
// 		err := setupTestDB()
// 		assert.Nil(t, err)

// 		credits := big.NewInt(1000)
// 		priceLimit := big.NewInt(500)

// 		auction := &entity.Auction{
// 			Credits:    credits,
// 			PriceLimit: priceLimit,
// 			State:      entity.AuctionOngoing,
// 			ExpiresAt:  time.Now().Add(24 * time.Hour).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 			UpdatedAt:  0,
// 		}
// 		res, err := repo.CreateAuction(auction)
// 		assert.Nil(t, err)

// 		err = repo.DeleteAuction(res.Id)
// 		assert.Nil(t, err)

// 		storedAuction, err := repo.FindAuctionById(res.Id)
// 		assert.NotNil(t, err)
// 		assert.Nil(t, storedAuction)
// 	})
// }
