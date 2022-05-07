package model

type SensorResponse struct {
	ID          string  `json:"id"`
	NodeID      string  `json:"nodeId"`
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temperature"`
	HeatIndex   float64 `json:"heatIndex"`
}

type GetAllSensorResponse struct {
	Items []*SensorResponse `json:"items"`
}

type CreateSensorRequest struct {
	Humidity    float64 `json:"humidity" validate:"required" label:"humidity"`
	Temperature float64 `json:"temperature" validate:"required" label:"temperature"`
	HeatIndex   float64 `json:"heatIndex" validate:"required" label:"heatIndex"`
}

type ShowSensorRequest struct {
	ID string `param:"id" validate:"required" label:"id"`
}

type UpdateSensorRequest struct {
	ID          string  `param:"id" validate:"required" label:"id"`
	Humidity    float64 `json:"humidity" validate:"required" label:"humidity"`
	Temperature float64 `json:"temperature" validate:"required" label:"temperature"`
	HeatIndex   float64 `json:"heatIndex" validate:"required" label:"heatIndex"`
}

type DeleteSensorRequest struct {
	ID string `param:"id" validate:"required" label:"id"`
}
