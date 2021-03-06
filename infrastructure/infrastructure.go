package infrastructure

import (
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/krobus00/iot-be/infrastructure/validator"
	kro_pkg "github.com/krobus00/krobot-building-block/pkg"
	"go.uber.org/fx"
)

type Infrastructure struct {
	fx.In

	Logger          Logger
	Translator      *ut.UniversalTranslator
	Router          Router
	Env             Env
	Database        kro_pkg.Database
	HttpClient      *http.Client
	Redis           *redis.Client
	CustomValidator validator.CustomValidator
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewTranslator),
	fx.Provide(NewRouter),
	fx.Provide(NewEnv),
	fx.Provide(NewHttpCleint),
	fx.Provide(NewDatabase),
	fx.Provide(NewRedisClient),
	fx.Provide(NewValidator),
	fx.Provide(validator.New),
	// fx.Invoke(RegisterCustomValidation),

	fx.Invoke(InitSentry),

	fx.Populate(NewInfrastructure()),
)
