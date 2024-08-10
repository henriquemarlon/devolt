package entity

import (
	"errors"
	"math/big"

	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidStation  = errors.New("invalid station")
	ErrStationNotFound = errors.New("station not found")
)

type StationRepository interface {
	CreateStation(station *Station) (*Station, error)
	FindStationById(id uint) (*Station, error)
	FindAllStations() ([]*Station, error)
	UpdateStation(station *Station) (*Station, error)
	DeleteStation(id uint) error
}

type StationState string

const (
	StationStateActive   StationState = "active"
	StationStateInactive StationState = "inactive"
)

type Station struct {
	Id             uint                `json:"id" gorm:"primaryKey"`
	Consumption    custom_type.BigInt  `json:"consumption,omitempty" gorm:"type:bigint;not null"`
	Owner          custom_type.Address `json:"owner,omitempty" gorm:"not null"`
	State          StationState        `json:"state,omitempty" gorm:"type:text;not null"`
	Orders         []*Order            `json:"orders,omitempty" gorm:"foreignKey:StationId;constraint:OnDelete:CASCADE"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit,omitempty" gorm:"type:bigint;not null"`
	Latitude       float64             `json:"latitude,omitempty" gorm:"not null"`
	Longitude      float64             `json:"longitude,omitempty" gorm:"not null"`
	CreatedAt      int64               `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt      int64               `json:"updated_at,omitempty" gorm:"default:0"`
}

func NewStation(owner custom_type.Address, pricePerCredit custom_type.BigInt, latitude float64, longitude float64, createdAt int64) (*Station, error) {
	station := &Station{
		Consumption:    custom_type.NewBigInt(big.NewInt(0)),
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
	if s.Owner.Address == (common.Address{}) || s.PricePerCredit.Int == nil || s.Latitude == 0 || s.Longitude == 0 || s.CreatedAt == 0 {
		return ErrInvalidStation
	}
	return nil
}
