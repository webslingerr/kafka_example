package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

// @Router /v1/brand/{id} [get]
// @Summary Get Brand
// @Description API for getting brand
// @Tags brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.BrandModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetBrand(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Error("Error whiling getting brand", logger.Any("error", err))
		h.HandleBadRequest(c, err, "Id format should be uuid")
		return
	}

	resp, err := h.storage.Brand().Get(id)
	if err != nil {
		h.log.Error("Error whiling getting brand", logger.Any("error", err))
		h.HandleError(c, err, "Brand not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /v1/brand [get]
// @Summary Get brands
// @Description API for get brands
// @Tags brand
// @Accept json
// @Produce json
// @Param find query response.GetAllBrandsRequest false "filters"
// @Success 201 {object} response.GetAllBrandsResponse
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllBrands(c *gin.Context) {
	var (
		queryParam car24_car_service.BrandQueryParamModel
		err        error
	)

	queryParam.Offset, err = ParseQueryParam(c, "offset", "0")
	if err != nil {
		h.HandleError(c, err, "wrong offset input")
		return
	}
	queryParam.Limit, err = ParseQueryParam(c, "limit", "10")
	if err != nil {
		h.HandleError(c, err, "wrong limit input")
		return
	}

	h.log.Info("----GETALL_BRANDSS--->", logger.Any("Request", queryParam))

	resp, err := h.storage.Brand().GetAll(queryParam)
	if err != nil {
		h.log.Error("Error whiling getting car", logger.Any("error", err))
		h.HandleInternalServerError(c, err, "Something went wrong")
		return
	}

	h.log.Info("----GETALL_BRANDS--->", logger.Any("Response", resp))

	h.handleSuccessResponse(c, 200, "ok", resp)
}
