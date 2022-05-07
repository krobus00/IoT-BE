package node

import (
	"context"
	"fmt"

	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/microcosm-cc/bluemonday"
)

func (svc *service) Register(ctx context.Context, payload *model.RegisterRequest) (*string, error) {
	span := kro_util.StartTracing(ctx, tag, tracingRegister)
	defer span.Finish()

	p := bluemonday.UGCPolicy()

	input := &db_models.Node{
		ID:        kro_util.NewUUID(),
		City:      p.Sanitize(payload.City),
		Longitude: payload.Longitude,
		Latitude:  payload.Latitude,
	}

	err := svc.repository.NodeRepository.Store(ctx, svc.db, input)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingRegister, err))
		return nil, err
	}

	return &input.ID, nil
}
