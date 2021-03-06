package model

type NodeResponse struct {
	ID        string  `json:"id"`
	City      string  `json:"city"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	DateColumn
}

type RegisterRequest struct {
	City      string  `json:"city" validate:"required" label:"city"`
	Longitude float64 `json:"longitude" validate:"required,uniquedb=nodes longitude" label:"longitude"`
	Latitude  float64 `json:"latitude" validate:"required,uniquedb=nodes latitude" label:"latitude"`
}

type GetAccessTokenRequest struct {
	ID string `json:"id" validate:"required,existdb=nodes node_id" label:"id"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GetNodeInfoPayload struct {
	ID string `param:"id" validate:"required,existdb=nodes node_id" label:"id"`
}

type GetNodeInfoResponse struct {
	NodeResponse
	LastReport *SensorResponse `json:"lastReport"`
}
