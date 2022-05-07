package sensor

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) UpdateSensorByID(ctx context.Context, payload *model.UpdateSensorRequest) error {
	span := kro_util.StartTracing(ctx, tag, tracingUpdateSensorByID)
	defer span.Finish()

	sensor, err := svc.GetSensorByID(ctx, &model.ShowSensorRequest{ID: payload.ID})
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingUpdateSensorByID, err))
		return err
	}

	if sensor == nil {
		return kro_model.NewHttpCustomError(http.StatusNotFound, errors.New("Sensor Not Found"))
	}

	input := &db_models.Sensor{
		ID:          payload.ID,
		NodeID:      sensor.NodeID,
		Humidity:    payload.Humidity,
		Temperature: payload.Temperature,
		HeatIndex:   payload.HeatIndex,
	}

	err = svc.repository.SensorRepository.UpdateSensorByID(ctx, svc.db, input)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingUpdateSensorByID, err))
		return err
	}

	return nil
}
