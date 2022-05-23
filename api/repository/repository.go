package repository

import (
	"github.com/krobus00/iot-be/api/repository/node"
	"github.com/krobus00/iot-be/api/repository/sensor"
	// IMPORT REPOSITORY PACKAGE HERE
	"go.uber.org/fx"
)

type Repository struct {
	fx.In

	NodeRepository   node.NodeRepository
	SensorRepository sensor.SensorRepository
	// INJECT REPOSITORY HERE
}

var Module = fx.Options(
	fx.Provide(node.New),
	fx.Provide(sensor.New),

	// INJECT REPOSITORY MODULE HERE
)
