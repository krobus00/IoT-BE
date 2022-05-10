package node

import (
	"context"
	"fmt"

	"github.com/krobus00/iot-be/model"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
)

func (svc *service) GetAllNodes(ctx context.Context, payload *kro_model.PaginationRequest) (*kro_model.PaginationResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetAllNodes)
	defer span.Finish()

	resp := new(kro_model.PaginationResponse)
	items := make([]*model.NodeResponse, 0)

	nodes, count, err := svc.repository.NodeRepository.GetAllNodes(ctx, svc.db, payload)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAllNodes, err))
		return nil, err
	}
	if nodes == nil {
		return nil, nil
	}
	for _, node := range nodes {
		items = append(items, &model.NodeResponse{
			ID:        node.ID,
			City:      node.City,
			Longitude: node.Longitude,
			Latitude:  node.Latitude,
			DateColumn: model.DateColumn{
				CreatedAt: node.CreatedAt,
				UpdatedAt: node.UpdatedAt,
				DeletedAt: node.DeletedAt,
			},
		})
	}
	resp.BuildResponse(payload, nodes, count)

	return resp, nil
}
