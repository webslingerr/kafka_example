package v1

import (
	"net/http"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/car24_car_service"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// @Security ApiKeyAuth
// @Router /v1/mark [post]
// @Summary Create Mark
// @Description API for creating mark
// @Tags mark
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param Mark body car24_car_service.CreateMarkModel true "mark"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) CreateMark(c *gin.Context) {
	var mark car24_car_service.CreateMarkModel

	err := c.ShouldBindJSON(&mark)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	id, _ := uuid.NewRandom()
	mark.ID = id.String()
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.mark.create")
	err = e.SetData(cloudevents.ApplicationJSON, mark)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.mark.create", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

// @Security ApiKeyAuth
// @Router /v1/mark/{id} [get]
// @Summary Get Mark
// @Description API for getting mark
// @Tags mark
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Success 200 {object} car24_car_service.MarkModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetMark(c *gin.Context) {
	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

// @Security ApiKeyAuth
// @Router /v1/mark [get]
// @Summary Get marks
// @Description API for getting all cars
// @Tags mark
// @Accept json
// @Produce json
// @Param find query car24_car_service.MarkQueryParamModel false "filters"
// @Success 200 {object} car24_car_service.MarkListModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllMarks(c *gin.Context) {

	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

// @Security ApiKeyAuth
// @Router /v1/mark/{id} [put]
// @Summary Update Mark
// @Description API for updating mark
// @Tags mark
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Param Mark body car24_car_service.UpdateMarkModel true "mark"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) UpdateMark(c *gin.Context) {
	var (
		mark car24_car_service.UpdateMarkModel
	)

	err := c.ShouldBindJSON(&mark)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	mark.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.mark.update")
	err = e.SetData(cloudevents.ApplicationJSON, mark)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.mark.update", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

// @Security ApiKeyAuth
// @Router /v1/mark/{id} [delete]
// @Summary Delete Mark
// @Description API for deleting mark
// @Tags mark
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) DeleteMark(c *gin.Context) {
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
	e.SetType("v1.car_service.mark.delete")
	err = e.SetData(cloudevents.ApplicationJSON, request)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.mark.delete", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}
