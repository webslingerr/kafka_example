package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

//@Router /v1/mark/{id} [get]
//@Summary Get Mark
//@Description API for getting mark
//@Tags mark
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} response.MarkModel
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) GetMark(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Error("Error whiling getting mark", logger.Any("error", err))
		h.HandleBadRequest(c, err, "Id format should be uuid")
		return
	}

	resp, err := h.storage.Mark().Get(id)
	if err != nil {
		h.log.Error("Error whiling getting mark", logger.Any("error", err))
		h.HandleError(c, err, "Mark not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

//@Router /v1/mark [get]
//@Summary Get marks
//@Description API for get marks
//@Tags mark
//@Accept json
//@Produce json
//@Param find query response.GetAllMarksRequest false "filters"
//@Success 201 {object} response.GetAllMarksResponse
//@Failure 400 {object} response.ResponseError
//@Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllMarks(c *gin.Context) {
	var (
		queryParam car24_car_service.MarkQueryParamModel
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

	h.log.Info("----GETALL_MARKS--->", logger.Any("Request", queryParam))

	resp, err := h.storage.Mark().GetAll(queryParam)
	if err != nil {
		h.log.Error("Error whiling getting mark", logger.Any("error", err))
		h.HandleInternalServerError(c, err, "Something went wrong")
		return
	}

	h.log.Info("----GETALL_MARKS--->", logger.Any("Response", resp))

	h.handleSuccessResponse(c, 200, "ok", resp)
}
