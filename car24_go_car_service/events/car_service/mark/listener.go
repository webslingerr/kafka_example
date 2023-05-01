package mark

import (
	"context"
	"encoding/json"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cs "gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

func (c *MarkService) Create(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		mark cs.CreateMarkModel
	)

	c.log.Info("Mark create", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &mark)

	resp, hasError := helper.HandleBadRequest(c.log, mark.ID, "Error while unmarshalling", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	_, err = c.storage.Mark().Create(&mark)

	resp, hasError = helper.HandleInternal(c.log, mark.ID, "Error while creating mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	c.log.Info("Successfully created Mark", logger.String("id", mark.ID))

	err = c.kafka.Push("v1.car_service.mark.created", event)
	resp, hasError = helper.HandleInternal(c.log, mark.ID, "Error while creating mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	resp = response.Response{
		ID:         mark.ID,
		StatusCode: http.StatusOK,
		Data:       mark,
		Action:     "create",
	}

	return resp
}

func (c *MarkService) Update(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		mark cs.UpdateMarkModel
	)

	c.log.Info("Mark update", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &mark)
	resp, hasError := helper.HandleBadRequest(c.log, mark.ID, "Error while marshaling", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	err = c.storage.Mark().Update(&mark)
	resp, hasError = helper.HandleInternal(c.log, mark.ID, "Error while updating mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	c.log.Info("Successfully Updated Mark", logger.String("id", mark.ID))

	err = c.kafka.Push("v1.car_service.mark.updated", event)
	resp, hasError = helper.HandleInternal(c.log, mark.ID, "Error while updating mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	return response.Response{
		ID:         mark.ID,
		StatusCode: http.StatusOK,
		Action:     "update",
	}
}

func (c *MarkService) Delete(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		mark cs.DeleteMarkModel
	)

	c.log.Info("Mark delete", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &mark)
	resp, hasError := helper.HandleInternal(c.log, mark.ID, "Error while unmarshalling mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	id := mark.ID

	err = c.storage.Mark().Delete(string(id))
	resp, hasError = helper.HandleInternal(c.log, mark.ID, "Error while deleting mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	c.log.Info("Successfully Deleted Mark", logger.String("id", mark.ID))

	err = c.kafka.Push("v1.car_service.mark.deleted", event)
	resp, hasError = helper.HandleInternal(c.log, mark.ID, "Error while deleting mark", err)
	if hasError {
		resp.CorrelationID = mark.ID
		return resp
	}

	return response.Response{
		ID:         mark.ID,
		StatusCode: http.StatusOK,
		Action:     "delete",
	}
}
