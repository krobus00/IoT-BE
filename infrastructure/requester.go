package infrastructure

import (
	"net/http"
	"time"
)

func NewHttpCleint(env Env) *http.Client {
	client := &http.Client{Timeout: time.Duration(env.HttpDefaultTimeOut) * time.Second}
	return client
}
