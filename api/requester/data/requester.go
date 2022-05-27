package data

import (
	"context"
	"net/http"

	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/model"
)

const (
	tag = `[DataRequester]`

	tracingResamplingData = "CallResamplingData"
	tracingForecastData   = "CallForecastData"
)

const (
	PROCESSING_DATA_ENDPOINT = "/_internal/data/preprocessing"
	FORECAST_DATA_ENDPOINT   = "/_internal/data/forecast"
)

type (
	Requester interface {
		CallResamplingData(context context.Context, payload *model.GetAllSensorResponse) ([]*model.GetSampledData, error)
		CallForecastData(context context.Context, payload *model.GetForecastRequest) ([]*model.GetForecastData, error)
	}
	requester struct {
		logger     infrastructure.Logger
		env        infrastructure.Env
		HttpClient *http.Client
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
) Requester {
	return &requester{
		logger:     infrastructure.Logger,
		env:        infrastructure.Env,
		HttpClient: infrastructure.HttpClient,
	}
}
