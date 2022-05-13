package node

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/krobus00/iot-be/api/service/node"
	"github.com/krobus00/iot-be/infrastructure"
	"go.uber.org/fx"
)

const (
	tag = "[NodeController]"

	tracingRegister        = "handleRegisterNode"
	tracingGetAccessToken  = "handleGetAccessTokenNode"
	tracingGetPagination   = "handleGetAccessTokenNode"
	tracingGetNodeInfo     = "handleGetNodeInfo"
	tracingGetNodeInfoByID = "handleGetNodeInfoByID"
)

type Controller struct {
	fx.In

	Logger      infrastructure.Logger
	Translator  *ut.UniversalTranslator
	NodeService node.Service
}
