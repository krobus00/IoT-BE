package sensor

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/model"
	db_models "github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
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
	tracingCountSensors          = "CountSensors"
)

type (
	SensorRepository interface {
		GetTableName() string
		GetAllSensor(ctx context.Context, db *sqlx.DB, paginationRequest *kro_model.PaginationRequest, config ...kro_model.Config) ([]*db_models.Sensor, error)
		Store(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) error
		GetSensorByID(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) (*db_models.Sensor, error)
		UpdateSensorByID(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) error
		DeleteSensorByID(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) error
		GetLastReportByNodeID(ctx context.Context, db kro_pkg.Querier, input *db_models.Sensor) (*db_models.Sensor, error)
		GetSensorByRange(ctx context.Context, db kro_pkg.Querier, input *model.GetProcessedDataRequest) ([]*db_models.Sensor, error)
		CountSensors(ctx context.Context, db kro_pkg.Querier) (int64, error)
	}
	repository struct {
		logger infrastructure.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) SensorRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetTableName() string {
	return "sensors"
}
