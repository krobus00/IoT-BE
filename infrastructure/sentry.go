package infrastructure

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry(env Env) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              env.SentryDSN,
		Environment:      fmt.Sprintf("[%s][%s]", env.AppName, env.Environment),
		AttachStacktrace: true,
		TracesSampleRate: 0.2,
		SampleRate:       0.2,
	})
	if err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}
	defer sentry.Recover()
}
