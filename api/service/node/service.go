package node

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/model"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

const (
	tag = `[NodeService]`

	tracingRegister       = "RegisterNode"
	tracingGetAccessToken = "GetAccessToken"
	tracingGetAllNodes    = "GetAllNodes"
	tracingGetNodeInfo    = "GetNodeInfo"
	tracingGetNodeByID    = "GetNodeByID"
)

type (
	NodeService interface {
		Register(ctx context.Context, payload *model.RegisterRequest) (*string, error)
		GetAccessToken(ctx context.Context, payload *model.GetAccessTokenRequest) (*model.GetAccessTokenResponse, error)
		GetAllNodes(ctx context.Context, payload *kro_model.PaginationRequest) (*kro_model.PaginationResponse, error)
		GetNodeInfo(ctx context.Context) (*model.GetNodeInfoResponse, error)
		GetNodeInfoByID(ctx context.Context, payload *model.GetNodeInfoPayload) (*model.GetNodeInfoResponse, error)
	}
	service struct {
		logger     infrastructure.Logger
		db         *sqlx.DB
		repository repository.Repository
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
	repository repository.Repository,
) NodeService {
	return &service{
		logger:     infrastructure.Logger,
		db:         infrastructure.Database.SqlxDB,
		repository: repository,
	}
}
