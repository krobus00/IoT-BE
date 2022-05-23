package requester

import (
	"github.com/krobus00/iot-be/api/requester/data"
	// IMPORT REQUESTER PACKAGE HERE
	"go.uber.org/fx"
)

type Requester struct {
	fx.In

	DataRequester data.Requester
	// INJECT REQUESTER HERE
}

var Module = fx.Options(
	fx.Provide(data.New),
	// INJECT REQUESTER MODULE HERE
)
