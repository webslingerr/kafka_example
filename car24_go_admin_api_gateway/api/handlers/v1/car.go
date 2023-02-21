package v1

import (
	"net/http"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/car24_car_service"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

//@Security ApiKeyAuth
//@Router /v1/car [post]
//@Summary Create Car
//@Description API for creating car
//@Tags car
//@Accept json
//@Produce json
//@Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
//@Param Car body car24_car_service.CreateCarModel true "car"
//@Success 201 {object} response.ResponseOK
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) CreateCar(c *gin.Context) {
	var car car24_car_service.CreateCarModel

	err := c.ShouldBindJSON(&car)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	id, _ := uuid.NewRandom()
	car.ID = id.String()
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.car.create")
	err = e.SetData(cloudevents.ApplicationJSON, car)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.car.create", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

//@Security ApiKeyAuth
//@Router /v1/car/{id} [get]
//@Summary Get Car
//@Description API for getting car
//@Tags car
//@Accept json
//@Produce json
//@Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
//@Param id path string true "id"
//@Success 200 {object} car24_car_service.CarModel
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) GetCar(c *gin.Context) {
	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

//@Security ApiKeyAuth
//@Router /v1/car [get]
//@Summary Get cars
//@Description API for getting all cars
//@Tags car
//@Accept json
//@Produce json
//@Param find query car24_car_service.CarQueryParamModel false "filters"
//@Success 200 {object} car24_car_service.CarListModel
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllCars(c *gin.Context) {

	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

//@Security ApiKeyAuth
//@Router /v1/car/{id} [put]
//@Summary Update Car
//@Description API for updating car
//@Tags car
//@Accept json
//@Produce json
//@Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
//@Param id path string true "id"
//@Param Car body car24_car_service.UpdateCarModel true "car"
//@Success 201 {object} response.ResponseOK
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) UpdateCar(c *gin.Context) {
	var (
		car car24_car_service.UpdateCarModel
	)

	err := c.ShouldBindJSON(&car)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	car.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.car.update")
	err = e.SetData(cloudevents.ApplicationJSON, car)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.car.update", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

//@Security ApiKeyAuth
//@Router /v1/car/{id} [delete]
//@Summary Delete Car
//@Description API for deleting car
//@Tags car
//@Accept json
//@Produce json
//@Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
//@Param id path string true "id"
//@Success 201 {object} response.ResponseOK
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) DeleteCar(c *gin.Context) {
	var (
		request car24_car_service.DeleteCarModel
	)

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	request.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.car.delete")
	err = e.SetData(cloudevents.ApplicationJSON, request)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.car.delete", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}
