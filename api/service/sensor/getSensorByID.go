package sensor

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetSensorByID(ctx context.Context, payload *model.ShowSensorRequest) (*model.SensorResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetSensorByID)
	defer span.Finish()

	sensor, err := svc.repository.SensorRepository.GetSensorByID(ctx, svc.db, &database.Sensor{ID: payload.ID})
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetSensorByID, err))
		return nil, err
	}
	if sensor == nil {
		return nil, kro_model.NewHttpCustomError(http.StatusNotFound, errors.New("Sensor Not Found"))
	}
	resp := &model.SensorResponse{
		ID:          sensor.ID,
		NodeID:      sensor.NodeID,
		Humidity:    sensor.Humidity,
		Temperature: sensor.Temperature,
		HeatIndex:   sensor.HeatIndex,
	}
	return resp, nil
}
