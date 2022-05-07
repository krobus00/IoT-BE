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

	tracingRegister       = "RegisterUser"
	tracingGetAccessToken = "GetAccessToken"
	tracingGetAllNodes    = "GetAllNodes"
)

type (
	Service interface {
		Register(ctx context.Context, payload *model.RegisterRequest) (*string, error)
		GetAccessToken(ctx context.Context, payload *model.GetAccessTokenRequest) (*model.GetAccessTokenResponse, error)
		GetAllNodes(ctx context.Context, payload *kro_model.PaginationRequest) (*kro_model.PaginationResponse, error)
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
) Service {
	return &service{
		logger:     infrastructure.Logger,
		db:         infrastructure.Database.SqlxDB,
		repository: repository,
	}
}
