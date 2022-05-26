package sensor

import (
	"context"
	"fmt"
	"time"

	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetForecastData(ctx context.Context, payload *model.GetForecastDataRequest) (*model.GetForecastDataResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetForecastData)
	defer span.Finish()

	sensorData := new(model.GetAllSensorResponse)

	// redisKey := fmt.Sprintf("FORECAST-%s-%s", payload.NodeID, time.Now().Format("2006-01-02"))

	payloadData := &model.GetProcessedDataRequest{
		NodeID:    payload.NodeID,
		StartDate: time.Now().Add(-time.Hour * 24).Unix(),
		EndDate:   time.Now().Unix(),
	}

	sensors, err := svc.repository.SensorRepository.GetSensorByRange(ctx, svc.db, payloadData)

	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetForecastData, err))
		return nil, err
	}
	if sensors == nil {
		return nil, nil
	}

	for _, sensor := range sensors {
		sensorData.Items = append(sensorData.Items, &model.SensorResponse{
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

	node, err := svc.repository.NodeRepository.FindNodeByID(ctx, svc.db, &database.Node{ID: payload.NodeID})

	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetForecastData, err))
		return nil, err
	}

	forecastPayload := &model.GetForecastRequest{
		Model: node.ModelURL,
		Items: sensorData.Items,
	}

	data, err := svc.requester.DataRequester.CallForecastData(ctx, forecastPayload)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetForecastData, err))
		return nil, err
	}

	resp := new(model.GetForecastDataResponse)
	for _, v := range data {
		resp.DateTime = append(resp.DateTime, &v.CreatedAt)
		resp.Temperature = append(resp.Temperature, &v.Temperature)
	}
	return resp, nil
}
