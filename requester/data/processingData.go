package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/krobus00/iot-be/model"
)

func (r *requester) CallResamplingData(context context.Context, payload *model.GetAllSensorResponse) ([]*model.GetSampledData, error) {
	jsonData, err := json.Marshal(payload)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s%s", r.env.DataServiceHost, PROCESSING_DATA_ENDPOINT),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingResamplingData, err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := r.HttpClient
	resp, err := client.Do(req)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingResamplingData, err))
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingResamplingData, err))
		return nil, err
	}
	res := make([]*model.GetSampledData, 0)
	if err := json.Unmarshal(responseBody, &res); err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with: %v", tag, tracingResamplingData, err))
		return nil, err
	}

	return res, nil
}
