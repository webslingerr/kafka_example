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
// @Router /v1/brand [post]
// @Summary Create Brand
// @Description API for creating brand
// @Tags brand
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param Car body car24_car_service.CreateBrandModel true "brand"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) CreateBrand(c *gin.Context) {
	var brand car24_car_service.CreateBrandModel

	err := c.ShouldBindJSON(&brand)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	id, _ := uuid.NewRandom()
	brand.ID = id.String()
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.brand.create")
	err = e.SetData(cloudevents.ApplicationJSON, brand)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.brand.create", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

// @Security ApiKeyAuth
// @Router /v1/brand/{id} [get]
// @Summary Get Brand
// @Description API for getting brand
// @Tags brand
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Success 200 {object} car24_car_service.BrandModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetBrand(c *gin.Context) {
	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

// @Security ApiKeyAuth
// @Router /v1/brand [get]
// @Summary Get brands
// @Description API for getting all brand
// @Tags brand
// @Accept json
// @Produce json
// @Param find query car24_car_service.BrandQueryParamModel false "filters"
// @Success 200 {object} car24_car_service.BrandListModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllBrands(c *gin.Context) {

	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

// @Security ApiKeyAuth
// @Router /v1/brand/{id} [put]
// @Summary Update Brand
// @Description API for updating brand
// @Tags brand
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Param Car body car24_car_service.UpdateBrandModel true "car"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) UpdateBrand(c *gin.Context) {
	var (
		brand car24_car_service.UpdateBrandModel
	)

	err := c.ShouldBindJSON(&brand)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	brand.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.brand.update")
	err = e.SetData(cloudevents.ApplicationJSON, brand)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.brand.update", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

// @Security ApiKeyAuth
// @Router /v1/brand/{id} [delete]
// @Summary Delete Brand
// @Description API for deleting brand
// @Tags brand
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) DeleteBrand(c *gin.Context) {
	var (
		request car24_car_service.DeleteBrandModel
	)

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	request.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.brand.delete")
	err = e.SetData(cloudevents.ApplicationJSON, request)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.brand.delete", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}
