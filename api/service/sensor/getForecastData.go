package sensor

import (
	"context"
	"fmt"
	"time"

	"github.com/krobus00/iot-be/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetForecastData(ctx context.Context, payload *model.GetForecastDataRequest) (*model.GetForecastDataResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetResampledData)
	defer span.Finish()

	items := new(model.GetAllSensorResponse)

	// redisKey := fmt.Sprintf("FORECAST-%s-%s", payload.NodeID, time.Now().Format("2006-01-02"))

	payloadData := &model.GetProcessedDataRequest{
		NodeID:    payload.NodeID,
		StartDate: model.DateTime(time.Now().Add(-time.Hour * 24)),
		EndDate:   model.DateTime(time.Now()),
	}

	sensors, err := svc.repository.SensorRepository.GetSensorByRange(ctx, svc.db, payloadData)

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

	data, err := svc.requester.DataRequester.CallForecastData(ctx, items)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetResampledData, err))
		return nil, err
	}

	resp := new(model.GetForecastDataResponse)
	for _, v := range data {
		resp.DateTime = append(resp.DateTime, &v.CreatedAt)
		resp.Temperature = append(resp.Temperature, &v.Temperature)
	}
	return resp, nil
}
