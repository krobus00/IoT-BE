package sensor

import (
	"context"
	"fmt"

	"github.com/krobus00/iot-be/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetResampledData(ctx context.Context, payload *model.GetProcessedDataRequest) (*model.GetProcessedDataResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetResampledData)
	defer span.Finish()

	items := new(model.GetAllSensorResponse)

	sensors, err := svc.repository.SensorRepository.GetSensorByRange(ctx, svc.db, payload)

	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetResampledData, err))
		return nil, err
	}
	if sensors == nil {
		return nil, nil
	}

	for _, sensor := range sensors {
		items.Items = append(items.Items, &model.SensorResponse{
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

	data, err := svc.requester.DataRequester.CallResamplingData(ctx, items)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetResampledData, err))
		return nil, err
	}

	resp := new(model.GetProcessedDataResponse)
	for _, v := range data {
		resp.DateTime = append(resp.DateTime, v.CreatedAt)
		resp.Humidity = append(resp.Humidity, v.Humidity)
		resp.Temperature = append(resp.Temperature, v.Temperature)
		resp.HeatIndex = append(resp.HeatIndex, v.HeatIndex)
	}
	return resp, nil
}
