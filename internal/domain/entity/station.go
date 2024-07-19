package entity

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type StationRepository interface {
	CreateStation(station *Station) (*Station, error)
	FindStationById(id string) (*Station, error)
	FindAllStations() ([]*Station, error)
	UpdateStation(station *Station) (*Station, error)
	DeleteStation(id string) error
}

type Station struct {
	Id             string         `json:"id" gorm:"primaryKey"`
	Consumption    *big.Int       `json:"consumption" gorm:"type:bigint"`
	Owner          common.Address `json:"owner" gorm:"not null"`
	State          string         `json:"state" gorm:"type:text;default:'pending'"`
	Orders         []*Order       `json:"orders" gorm:"foreignKey:StationId;constraint:OnDelete:CASCADE"`
	PricePerCredit *big.Int       `json:"price_per_credit" gorm:"type:bigint;not null"`
	Latitude       float64        `json:"latitude" gorm:"not null"`
	Longitude      float64        `json:"longitude" gorm:"not null"`
	CreatedAt      int64          `json:"created_at" gorm:"not null"`
	UpdatedAt      int64          `json:"updated_at" gorm:"default:0"`
}

func NewStation(id string, owner common.Address, pricePerCredit *big.Int, latitude float64, longitude float64, state string, createdAt int64) *Station {
	return &Station{
		Id:             id,
		Owner:          owner,
		State:          state,
		PricePerCredit: pricePerCredit,
		Latitude:       latitude,
		Longitude:      longitude,
		CreatedAt:      createdAt,
	}
}
