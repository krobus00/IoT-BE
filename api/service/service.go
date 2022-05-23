package service

import (
	"github.com/krobus00/iot-be/api/service/node"
	"github.com/krobus00/iot-be/api/service/sensor"
	// IMPORT SERVICE PACKAGE HERE
	"go.uber.org/fx"
)

type Service struct {
	fx.In

	NodeService   node.NodeService
	SensorService sensor.SensorService
	// INJECT SERVICE HERE
}

var Module = fx.Options(
	fx.Provide(node.New),
	fx.Provide(sensor.New),

// INJECT SERVICE MODULE HERE
)
