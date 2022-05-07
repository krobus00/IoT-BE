package controller

import (
	"github.com/krobus00/iot-be/api/controller/node"
	"github.com/krobus00/iot-be/api/controller/sensor"
	// IMPORT CONTROLLER PACKAGE HERE
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	NodeController   node.Controller
	SensorController sensor.Controller
	// INJECT CONTROLLER HERE
}

func NewHandler() *Handler {
	return &Handler{}
}

var Module = fx.Options(
	fx.Populate(NewHandler()),
)
