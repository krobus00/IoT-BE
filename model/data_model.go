package model

type GetSampledData struct {
	Humidity    float64  `json:"humidity"`
	Temperature float64  `json:"temperature"`
	HeatIndex   float64  `json:"heatIndex"`
	CreatedAt   DateTime `json:"createdAt"`
}

// type GetSampledDataResponse []*GetSampledData
