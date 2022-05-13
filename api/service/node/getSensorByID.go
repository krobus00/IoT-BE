package node

import (
	"context"
	"fmt"

	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetNodeInfoByID(ctx context.Context, payload *model.GetNodeInfoPayload) (*model.GetNodeInfoResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetNodeByID)
	defer span.Finish()

	resp := new(model.GetNodeInfoResponse)

	nodePayload := &db_models.Node{
		ID: payload.ID,
	}

	node, err := svc.repository.NodeRepository.FindNodeByID(ctx, svc.db, nodePayload)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetNodeInfo, err))
		return nil, err
	}
	if node == nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetNodeInfo, err))
		return nil, err
	}

	sensorPayload := &db_models.Sensor{
		NodeID: node.ID,
	}
	lastReport, err := svc.repository.SensorRepository.GetLastReportByNodeID(ctx, svc.db, sensorPayload)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetNodeInfo, err))
		return nil, err
	}
	resp.NodeResponse = model.NodeResponse{
		ID:        node.ID,
		City:      node.City,
		Longitude: node.Longitude,
		Latitude:  node.Latitude,
		DateColumn: model.DateColumn{
			CreatedAt: node.CreatedAt,
			UpdatedAt: node.UpdatedAt,
			DeletedAt: node.DeletedAt,
		},
	}
	resp.LastReport = &model.SensorResponse{
		ID:          lastReport.ID,
		NodeID:      node.ID,
		Humidity:    lastReport.Humidity,
		Temperature: lastReport.Temperature,
		HeatIndex:   lastReport.HeatIndex,

		DateColumn: model.DateColumn{
			CreatedAt: lastReport.CreatedAt,
			UpdatedAt: lastReport.UpdatedAt,
			DeletedAt: lastReport.DeletedAt,
		},
	}

	return resp, nil
}
