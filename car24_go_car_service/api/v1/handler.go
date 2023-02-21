package v1

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

type handlerV1 struct {
	log     logger.Logger
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:     options.Log,
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func (h *handlerV1) HandleError(c *gin.Context, err error, message string) {
	if err == sql.ErrNoRows {
		h.HandleNotFoundError(c, err, message)
	} else {
		h.HandleInternalServerError(c, err, message)
	}
}

func (h *handlerV1) HandleBadRequest(c *gin.Context, err error, message string) {
	h.log.Error(message, logger.Error(err))
	c.JSON(http.StatusNotFound, response.Response{
		Error: response.Error{
			Code:    helper.ErrorCodeBadRequest,
			Message: message,
		},
	})
}

func ParseQueryParam(c *gin.Context, key string, defaultValue string) (int, error) {
	valueStr := c.DefaultQuery(key, defaultValue)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return 0, err
	}

	return value, nil
}

func ParseStringToInt(q url.Values, key string) (value int, err error) {
	valueStr := q.Get(key)
	if valueStr == "" {
		valueStr = "0"
	}

	value, err = strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func GetLanguage(c *gin.Context) string {
	value := c.GetHeader("Accept-Language")
	if value == "" || value != "uz" && value != "ru" && value != "en" {
		value = "ru"
	}

	return value
}

func (h *handlerV1) HandleInternalServerError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, response.Response{
		Error: response.Error{
			Code:    helper.ErrorCodeBadRequest,
			Message: message,
		},
	})
}

func (h *handlerV1) HandleNotFoundError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusNotFound, response.Response{
		Error: response.Error{
			Code:    helper.ErrorCodeBadRequest,
			Message: message,
		},
	})
}

func (h *handlerV1) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, response.SuccessModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Data:    data,
	})
}
func (h *handlerV1) parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Query("offset"))
}

func (h *handlerV1) parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Query("limit"))
}
