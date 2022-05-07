package node

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_util "github.com/krobus00/krobot-building-block/util"
	"github.com/microcosm-cc/bluemonday"
)

func (svc *service) GetAccessToken(ctx context.Context, payload *model.GetAccessTokenRequest) (*model.GetAccessTokenResponse, error) {
	span := kro_util.StartTracing(ctx, tag, tracingGetAccessToken)
	defer span.Finish()

	p := bluemonday.UGCPolicy()

	input := &db_models.Node{
		ID: p.Sanitize(payload.ID),
	}

	node, err := svc.repository.NodeRepository.FindNodeByID(ctx, svc.db, input)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAccessToken, err))
		return nil, err
	}

	if node == nil {
		return nil, kro_model.NewHttpCustomError(http.StatusNotFound, errors.New("Incorrect node id"))
	}

	token, err := kro_util.GeneratePermanentToken(node.ID, "secret")
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingGetAccessToken, err))
		return nil, err
	}
	fmt.Println()
	fmt.Println(token)
	fmt.Println()

	resp := &model.GetAccessTokenResponse{
		AccessToken: token,
	}

	return resp, nil
}
