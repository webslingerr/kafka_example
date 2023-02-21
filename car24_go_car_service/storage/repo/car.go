package repo

import (
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
)

type CarI interface {
	Create(car *car24_car_service.CreateCarModel) (string, error)
	Get(id string) (*car24_car_service.CarModel, error)
	GetAll(car24_car_service.CarQueryParamModel) (car24_car_service.CarListModel, error)
	Update(car *car24_car_service.UpdateCarModel) error
	Delete(id string) error
}
