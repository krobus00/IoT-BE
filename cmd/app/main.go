package main

import (
	"github.com/joho/godotenv"
	"github.com/krobus00/iot-be/bootstrap"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.AppModule).Run()
}
