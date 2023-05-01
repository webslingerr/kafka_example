package brand

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

func (c *BrandService) Create(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		brand cs.CreateBrandModel
	)

	c.log.Info("Brand create", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &brand)

	resp, hasError := helper.HandleBadRequest(c.log, brand.ID, "Error while unmarshalling", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	_, err = c.storage.Brand().Create(&brand)

	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while creating brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	c.log.Info("Successfully created Brand", logger.String("id", brand.ID))

	err = c.kafka.Push("v1.car_service.brand.created", event)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while creating brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	resp = response.Response{
		ID:         brand.ID,
		StatusCode: http.StatusOK,
		Data:       brand,
		Action:     "create",
	}

	return resp
}

func (c *BrandService) Update(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		brand cs.UpdateBrandModel
	)

	c.log.Info("Brand update", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &brand)
	resp, hasError := helper.HandleBadRequest(c.log, brand.ID, "Error while marshaling", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	err = c.storage.Brand().Update(&brand)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while updating brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	c.log.Info("Successfully Updated Brand", logger.String("id", brand.ID))

	err = c.kafka.Push("v1.car_service.brand.updated", event)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while updating brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	return response.Response{
		ID:         brand.ID,
		StatusCode: http.StatusOK,
		Action:     "update",
	}
}

func (c *BrandService) Delete(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		brand cs.DeleteBrandModel
	)

	c.log.Info("Brand delete", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &brand)
	resp, hasError := helper.HandleInternal(c.log, brand.ID, "Error while unmarshalling brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	id := brand.ID

	err = c.storage.Brand().Delete(string(id))
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while deleting brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	c.log.Info("Successfully Deleted Brand", logger.String("id", brand.ID))

	err = c.kafka.Push("v1.car_service.brand.deleted", event)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while deleting brand", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	return response.Response{
		ID:         brand.ID,
		StatusCode: http.StatusOK,
		Action:     "delete",
	}
}
