package node

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

const (
	tag = `[NodeRepository]`

	tracingStore        = "StoreNode"
	tracingFindNodeByID = "FindNodeByID"
	tracingGetAllNodes  = "GetAllNodes"
	tracingCountNodes   = "CountNodes"
)

type (
	NodeRepository interface {
		GetTableName() string
		Store(ctx context.Context, db *sqlx.DB, input *db_models.Node) error
		FindNodeByID(ctx context.Context, db *sqlx.DB, input *db_models.Node) (*db_models.Node, error)
		GetAllNodes(ctx context.Context, db *sqlx.DB, paginationRequest *kro_model.PaginationRequest, config ...kro_model.Config) ([]*db_models.Node, error)
		CountNodes(ctx context.Context, db *sqlx.DB) (int64, error)
	}
	repository struct {
		logger infrastructure.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) NodeRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetTableName() string {
	return "nodes"
}
