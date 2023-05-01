package repo

import (
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
)

type MarkI interface {
	Create(car *car24_car_service.CreateMarkModel) (string, error)
	Get(id string) (*car24_car_service.MarkModel, error)
	GetAll(car24_car_service.MarkQueryParamModel) (car24_car_service.MarkListModel, error)
	Update(car *car24_car_service.UpdateMarkModel) error
	Delete(id string) error
}
