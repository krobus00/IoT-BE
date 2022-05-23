package sensor

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/api/requester"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/model"
	kro_model "github.com/krobus00/krobot-building-block/model"
)

const (
	tag = `[SensorService]`

	tracingGetAllSensor     = "GetAllSensor"
	tracingGetSensorByID    = "GetSensorByID"
	tracingStoreSensor      = "StoreSensor"
	tracingUpdateSensorByID = "UpdateSensorByID"
	tracingDeleteSensorByID = "DeleteSensorByID"
	tracingGetResampledData = "GetResampledData"
	tracingGetForecastData  = "GetForecastData"
)

type (
	SensorService interface {
		GetAllSensor(ctx context.Context, payload *kro_model.PaginationRequest) (*kro_model.PaginationResponse, error)
		StoreSensor(ctx context.Context, payload *model.CreateSensorRequest) error
		GetSensorByID(ctx context.Context, payload *model.ShowSensorRequest) (*model.SensorResponse, error)
		UpdateSensorByID(ctx context.Context, payload *model.UpdateSensorRequest) error
		DeleteSensorByID(ctx context.Context, payload *model.DeleteSensorRequest) error
		GetResampledData(ctx context.Context, payload *model.GetProcessedDataRequest) (*model.GetProcessedDataResponse, error)
		GetForecastData(ctx context.Context, payload *model.GetForecastDataRequest) (*model.GetForecastDataResponse, error)
	}
	service struct {
		logger     infrastructure.Logger
		db         *sqlx.DB
		repository repository.Repository
		requester  requester.Requester
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
	repository repository.Repository,
	requester requester.Requester,
) SensorService {
	return &service{
		logger:     infrastructure.Logger,
		db:         infrastructure.Database.SqlxDB,
		repository: repository,
		requester:  requester,
	}
}
