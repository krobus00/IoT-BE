package sensor

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

const (
	tag = `[SensorRepository]`

	tracingGetAllSensor          = "GetAllSensor"
	tracingGetSensorByID         = "GetSensorByID"
	tracingStoreSensor           = "StoreSensor"
	tracingUpdateSensorByID      = "UpdateSensorByID"
	tracingDeleteSensorByID      = "DeleteSensorByID"
	tracingGetLastReportByNodeID = "GetLastReportByNodeID"
	tracingGetSensorByRange      = "GetSensorByRange"
)

type (
	Repository interface {
		GetTableName() string
		GetAllSensor(ctx context.Context, db *sqlx.DB, paginationRequest *kro_model.PaginationRequest, config ...kro_model.Config) ([]*db_models.Sensor, int64, error)
		Store(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) error
		GetSensorByID(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) (*db_models.Sensor, error)
		UpdateSensorByID(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) error
		DeleteSensorByID(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) error
		GetLastReportByNodeID(ctx context.Context, db *sqlx.DB, input *db_models.Sensor) (*db_models.Sensor, error)
		GetSensorByRange(ctx context.Context, db *sqlx.DB, input *model.GetProcessedDataRequest) ([]*db_models.Sensor, error)
	}
	repository struct {
		logger infrastructure.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) Repository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetTableName() string {
	return "sensors"
}
