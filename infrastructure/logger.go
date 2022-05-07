package infrastructure

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(env Env) Logger {

	var config = zap.Config{}
	if env.Environment != "local" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build(zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level == zapcore.ErrorLevel {
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelError)
			})
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureMessage(fmt.Sprintf("%s:%d - %s \n %s", entry.Caller.File, entry.Caller.Line, entry.Message, entry))

		}
		if entry.Level == zapcore.WarnLevel {
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelWarning)
			})
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureMessage(fmt.Sprintf("%s:%d - %s \n %s", entry.Caller.File, entry.Caller.Line, entry.Message, entry))

		}
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.Clear()
		})
		return nil
	}))

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}

}
