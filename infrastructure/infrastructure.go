package infrastructure

import (
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
)

type Infrastructure struct {
	fx.In

	Logger     Logger
	Translator *ut.UniversalTranslator
	Router     Router
	Env        Env
	Database   Database
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewTranslator),
	fx.Provide(NewRouter),
	fx.Provide(NewEnv),
	fx.Provide(NewDatabase),
	fx.Provide(NewValidator),

	fx.Invoke(InitSentry),

	fx.Populate(NewInfrastructure()),
)
