package helper

import (
	"database/sql"
	"strconv"
	"strings"

	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

const (
	ErrorCodeInvalidURL    = "INVALID_URL"
	ErrorCodeInvalidJSON   = "INVALID_JSON"
	ErrorCodeInternal      = "INTERNAL"
	ErrorCodeUnauthorized  = "UNAUTHORIZED"
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	ErrorCodeNotFound      = "NOT_FOUND"
	ErrorCodeInvalidCode   = "INVALID_CODE"
	ErrorCodeBadRequest    = "BAD_REQUEST"
	ErrorCodeForbidden     = "FORBIDDEN"
	ErrorCodeNotApproved   = "NOT_APPROVED"
)

func HandleDBError(log logger.Logger, err error, message string, req interface{}) (e response.Response) {
	if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return response.Response{
			Error: response.Error{
				Code:    ErrorCodeNotFound,
				Message: message,
			},
		}
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return response.Response{
			Error: response.Error{
				Code:    ErrorCodeInternal,
				Message: message,
			},
		}

	}
	return
}
func HandleUnmarshallingError(log logger.Logger, sessionID string, err error, message string, req interface{}) response.Response {
	log.Error(message+", Bad Request", logger.Error(err), logger.Any("req", req))
	return response.Response{
		SessionID: sessionID,
		Error: response.Error{
			Code:    ErrorCodeBadRequest,
			Message: message,
		},
	}
}

func HandleBadRequest(log logger.Logger, sessionID, message string, err error) (response.Response, bool) {
	if err != nil {
		log.Error(message+", Bad Request", logger.Error(err))
		return response.Response{
			SessionID: sessionID,
			Error: response.Error{
				Code:    ErrorCodeBadRequest,
				Message: message,
			},
		}, true
	}

	return response.Response{}, false
}

func HandleInternal(log logger.Logger, sessionID, message string, err error) (response.Response, bool) {
	if err != nil {
		log.Error(message, logger.Error(err))
		return response.Response{
			SessionID: sessionID,
			Error: response.Error{
				Code:    ErrorCodeInternal,
				Message: message,
			},
		}, true
	}

	return response.Response{}, false
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
