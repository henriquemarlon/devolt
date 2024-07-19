package sqlite

import (
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"gorm.io/gorm"
)

type StationRepositorySqlite struct {
	Db *gorm.DB
}

func NewStationRepositorySqlite(db *gorm.DB) *StationRepositorySqlite {
	return &StationRepositorySqlite{
		Db: db,
	}
}

func (r *StationRepositorySqlite) CreateStation(input *entity.Station) (*entity.Station, error) {
	err := r.Db.Create(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create station: %w", err)
	}
	return input, nil
}

func (r *StationRepositorySqlite) FindStationById(id string) (*entity.Station, error) {
	var station entity.Station
	err := r.Db.Preload("Bids").First(&station, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find station by ID: %w", err)
	}
	return &station, nil
}

func (r *StationRepositorySqlite) FindAllStations() ([]*entity.Station, error) {
	var stations []*entity.Station
	err := r.Db.Preload("Bids").Find(&stations).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find all stations: %w", err)
	}
	return stations, nil
}

func (r *StationRepositorySqlite) UpdateStation(input *entity.Station) (*entity.Station, error) {
	err := r.Db.Save(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update station: %w", err)
	}
	return input, nil
}

func (r *StationRepositorySqlite) DeleteStation(id string) error {
	err := r.Db.Delete(&entity.Station{}, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("failed to delete station: %w", err)
	}
	return nil
}
