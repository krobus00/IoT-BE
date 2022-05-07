package sensor

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/krobus00/iot-be/api/service/sensor"
	"github.com/krobus00/iot-be/infrastructure"
	"go.uber.org/fx"
)

const (
	tag = "[SensorController]"

	tracingGetAllSensor = "handleGetAllSensor"
	tracingShowSensor   = "handleShowSensor"
	tracingStoreSensor  = "handleCreateSensor"
	tracingUpdateSensor = "handleUpdateSensor"
	tracingDeleteSensor = "handleDeleteSensor"
)

type Controller struct {
	fx.In

	Logger        infrastructure.Logger
	Translator    *ut.UniversalTranslator
	SensorService sensor.Service
}
