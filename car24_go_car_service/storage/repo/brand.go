package repo

import (
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
)

type BrandI interface {
	Create(car *car24_car_service.CreateBrandModel) (string, error)
	Get(id string) (*car24_car_service.BrandModel, error)
	GetAll(car24_car_service.BrandQueryParamModel) (car24_car_service.BrandListModel, error)
	Update(car *car24_car_service.UpdateBrandModel) error
	Delete(id string) error
}
