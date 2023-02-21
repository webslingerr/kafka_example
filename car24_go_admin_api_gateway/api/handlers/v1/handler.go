package v1

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/config"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/models"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/event"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/logger"
)

type handlerV1 struct {
	log   logger.Logger
	cfg   config.Config
	kafka *event.Kafka
}

//HandlerV1Config ...
type HandlerV1Config struct {
	Logger logger.Logger
	Cfg    config.Config
	Kafka  *event.Kafka
}

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorCodeBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
)

//New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:   c.Logger,
		cfg:   c.Cfg,
		kafka: c.Kafka,
	}
}

func (h *handlerV1) HandleBadRequest(c *gin.Context, err error, message string) {
	c.JSON(http.StatusBadRequest, models.ResponseError{
		Error: models.Error{
			Code:    ErrorCodeBadRequest,
			Message: message,
		},
	})
}

func (h *handlerV1) HandleInternalServerError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, models.ResponseError{
		Error: models.Error{
			Code:    ErrorCodeInternal,
			Message: message,
		},
	})
}

func (h *handlerV1) HandleBadRequestErrWithMessage(c *gin.Context, err error, message string) {
	h.log.Error(message, logger.Error(err))
	c.JSON(http.StatusBadRequest, models.ResponseError{
		Error: models.Error{
			Code:    ErrorCodeBadRequest,
			Message: message,
		},
	})
}

func HandleError(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.Error{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func (h *handlerV1) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, response.ErrorModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Error:   err,
	})
}

func (h *handlerV1) MakeProxy(c *gin.Context, proxyUrl, path string) (err error) {
	req := c.Request

	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		h.log.Error("error in parse addr: %v", logger.Error(err))
		c.String(http.StatusInternalServerError, "error")
		return
	}

	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	req.URL.Path = path
	transport := http.DefaultTransport
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// req = req.WithContext(ctx)
	resp, err := transport.RoundTrip(req)
	if HandleError(c, h.log, err, "error in round trip:"+proxyUrl+path) {
		return
	}

	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	_, _ = bufio.NewReader(resp.Body).WriteTo(c.Writer)
	return
}

func (h *handlerV1) MakeProxyValidator(c *gin.Context, proxyUrl, path, searchKey, searchValue string) (err error) {
	req := c.Request

	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		h.log.Error("error in parse addr: %v", logger.Error(err))
		c.String(http.StatusInternalServerError, "error")
		return
	}

	req.URL.Path = path
	req.Method = "GET"
	req.URL.RawQuery = fmt.Sprintf("%v=%v", searchKey, searchValue)

	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	transport := http.DefaultTransport

	resp, err := transport.RoundTrip(req)

	if HandleError(c, h.log, err, "error in round trip:"+proxyUrl+path) {
		return
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)

	return
}

func (h *handlerV1) GetSessionID(c *gin.Context) string {
	res, ok := c.Get("auth")
	if !ok {
		return ""
	}

	v := reflect.ValueOf(res)
	f := v.FieldByName("ID")
	return f.String()
}

func (h *handlerV1) GetCompanyID(c *gin.Context) string {
	res, ok := c.Get("auth")
	if !ok {
		return ""
	}

	v := reflect.ValueOf(res)
	f := v.FieldByName("CompanyID")
	return f.String()
}
