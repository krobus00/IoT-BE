package infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppName     string
	WorkerName  string
	AppPort     string
	WorkerPort  string
	Environment string

	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	SentryDSN             string
	SentryTraceSampleRate string
	SentrySampleRate      string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string

	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQVHost    string
}

func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

func (env *Env) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.AppName = os.Getenv("APP_NAME")
	env.AppPort = os.Getenv("APP_PORT")
	env.Environment = os.Getenv("ENVIRONMENT")

	env.DBUsername = os.Getenv("DATABASE_USERNAME")
	env.DBPassword = os.Getenv("DATABASE_PASSWORD")
	env.DBHost = os.Getenv("DATABASE_HOST")
	env.DBPort = os.Getenv("DATABASE_PORT")
	env.DBName = os.Getenv("DATABASE_NAME")

	env.SentryDSN = os.Getenv("SENTRY_DSN")
	env.SentryTraceSampleRate = os.Getenv("SENTRY_TRACE_SAMPLE_RATE")
	env.SentrySampleRate = os.Getenv("SENTRY_SAMPLE_RATE")
}
