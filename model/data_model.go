package model

type GetSampledData struct {
	Humidity    float64  `json:"humidity"`
	Temperature float64  `json:"temperature"`
	HeatIndex   float64  `json:"heatIndex"`
	CreatedAt   DateTime `json:"createdAt"`
}

type GetForecastData struct {
	Temperature float64  `json:"yhat1"`
	CreatedAt   DateTime `json:"ds"`
}
