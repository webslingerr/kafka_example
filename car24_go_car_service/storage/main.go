package storage

import (
	"gitlab.udevs.io/car24/car24_go_car_service/storage/postgres"
	"gitlab.udevs.io/car24/car24_go_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Car() repo.CarI
}

type storagePg struct {
	carRepo repo.CarI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		carRepo: postgres.NewCar(db),
	}
}

func (s *storagePg) Car() repo.CarI {
	return s.carRepo
}
