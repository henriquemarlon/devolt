package entity

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidStation  = errors.New("invalid station")
	ErrStationNotFound = errors.New("station not found")
)

type StationRepository interface {
	CreateStation(station *Station) (*Station, error)
	FindStationById(id string) (*Station, error)
	FindAllStations() ([]*Station, error)
	UpdateStation(station *Station) (*Station, error)
	DeleteStation(id string) error
}

type StationState string

const (
	StationStateActive   StationState = "active"
	StationStateInactive StationState = "inactive"
)

type Station struct {
	Id             string         `json:"id" gorm:"primaryKey"`
	Consumption    *big.Int       `json:"consumption" gorm:"type:bigint"`
	Owner          common.Address `json:"owner" gorm:"not null"`
	State          StationState   `json:"state" gorm:"type:text;not null"`
	Orders         []*Order       `json:"orders" gorm:"foreignKey:StationId;constraint:OnDelete:CASCADE"`
	PricePerCredit *big.Int       `json:"price_per_credit" gorm:"type:bigint;not null"`
	Latitude       float64        `json:"latitude" gorm:"not null"`
	Longitude      float64        `json:"longitude" gorm:"not null"`
	CreatedAt      int64          `json:"created_at" gorm:"not null"`
	UpdatedAt      int64          `json:"updated_at" gorm:"default:0"`
}

func NewStation(id string, owner common.Address, pricePerCredit *big.Int, latitude float64, longitude float64, createdAt int64) (*Station, error) {
	station := &Station{
		Id:             id,
		Owner:          owner,
		PricePerCredit: pricePerCredit,
		State:          StationStateActive,
		Latitude:       latitude,
		Longitude:      longitude,
		CreatedAt:      createdAt,
	}
	if err := station.Validate(); err != nil {
		return nil, err
	}
	return station, nil
}

func (s *Station) Validate() error {
	if s.Id == "" || s.Owner == (common.Address{}) || s.PricePerCredit == nil || s.Latitude == 0 || s.Longitude == 0 || s.CreatedAt == 0 {
		return ErrInvalidStation
	}
	return nil
}
