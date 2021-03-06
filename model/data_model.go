package model

type GetSampledData struct {
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temperature"`
	HeatIndex   float64 `json:"heatIndex"`
	CreatedAt   int64   `json:"createdAt"`
}

type GetForecastRequest struct {
	Model string            `json:"model"`
	Items []*SensorResponse `json:"items"`
}

type GetForecastData struct {
	Temperature float64 `json:"yhat1"`
	CreatedAt   int64   `json:"ds"`
}
