package sensor

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
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
)

type (
	Service interface {
		GetAllSensor(ctx context.Context, payload *kro_model.PaginationRequest) (*kro_model.PaginationResponse, error)
		StoreSensor(ctx context.Context, payload *model.CreateSensorRequest) error
		GetSensorByID(ctx context.Context, payload *model.ShowSensorRequest) (*model.SensorResponse, error)
		UpdateSensorByID(ctx context.Context, payload *model.UpdateSensorRequest) error
		DeleteSensorByID(ctx context.Context, payload *model.DeleteSensorRequest) error
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
