package sensor

import (
	"context"
	"fmt"

	"github.com/krobus00/iot-be/model"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetAllSensor(ctx context.Context, payload *kro_model.PaginationRequest) (*kro_model.PaginationResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetAllSensor)
	defer span.Finish()

	resp := new(kro_model.PaginationResponse)
	items := make([]*model.SensorResponse, 0)

	sensors, err := svc.repository.SensorRepository.GetAllSensor(ctx, svc.db, payload)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllSensor, err))
		return nil, err
	}
	if sensors == nil {
		return nil, nil
	}
	count, err := svc.repository.SensorRepository.CountSensors(ctx, svc.db)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllSensor, err))
		return nil, err
	}
	for _, sensor := range sensors {
		items = append(items, &model.SensorResponse{
			ID:          sensor.ID,
			NodeID:      sensor.NodeID,
			Humidity:    sensor.Humidity,
			Temperature: sensor.Temperature,
			HeatIndex:   sensor.HeatIndex,
			DateColumn: model.DateColumn{
				CreatedAt: sensor.CreatedAt,
				UpdatedAt: sensor.UpdatedAt,
				DeletedAt: sensor.DeletedAt,
			},
		})
	}
	resp.BuildResponse(payload, items, count)

	return resp, nil
}
