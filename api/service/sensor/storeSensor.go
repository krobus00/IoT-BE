package sensor

import (
	"context"
	"fmt"

	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) StoreSensor(ctx context.Context, payload *model.CreateSensorRequest) error {
	span := kro_util.StartTracing(ctx, tag, tracingStoreSensor)
	defer span.Finish()

	input := &db_models.Sensor{
		ID:          kro_util.NewUUID(),
		NodeID:      ctx.Value("nodeId").(string),
		Humidity:    payload.Humidity,
		Temperature: payload.Temperature,
		HeatIndex:   payload.HeatIndex,
	}

	err := svc.repository.SensorRepository.Store(ctx, svc.db, input)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStoreSensor, err))
		return err
	}

	return nil
}
