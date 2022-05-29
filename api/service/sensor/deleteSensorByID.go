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

func (svc *service) DeleteSensorByID(ctx context.Context, payload *model.DeleteSensorRequest) error {
	span := kro_util.StartTracing(ctx, tag, tracingDeleteSensorByID)
	defer span.Finish()

	sensor, err := svc.repository.SensorRepository.GetSensorByID(ctx, svc.db, &db_models.Sensor{ID: payload.ID})
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetSensorByID, err))
		return err
	}
	if sensor == nil {
		return kro_model.NewHttpCustomError(http.StatusNotFound, errors.New("Sensor Not Found"))
	}

	input := &db_models.Sensor{
		ID: sensor.ID,
	}

	err = svc.repository.SensorRepository.DeleteSensorByID(ctx, svc.db, input)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingDeleteSensorByID, err))
		return err
	}

	return nil
}
