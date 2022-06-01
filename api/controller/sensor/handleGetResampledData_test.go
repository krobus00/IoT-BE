package sensor

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/krobot-building-block/pkg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_HandleGetResampledData(t *testing.T) {
	type mockGetResampledData struct {
		res *model.GetProcessedDataResponse
		err error
	}
	type mockCustomValidator struct {
		res bool
	}
	tests := []struct {
		name                 string
		payload              string
		mockCustomValidator  *mockCustomValidator
		mockGetResampledData *mockGetResampledData
		wantStatusCode       int
		wantErr              bool
	}{
		{
			name: "WHEN get forecast sensor data data THEN system return success",
			payload: `{
				"nodeId": "UUID"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetResampledData: &mockGetResampledData{
				res: &model.GetProcessedDataResponse{},
				err: nil,
			},
			wantStatusCode: 200,
			wantErr:        false,
		},
		{
			name:    "WHEN get forecast sensor data data AND payload malformated THEN system return error",
			payload: `{`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name: "WHEN get forecast sensor data data AND payload invalid THEN system return error",
			payload: `{
				"nodeId": ""
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: false,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name: "WHEN get forecast sensor data data THEN system return error",
			payload: `{
				"id": "UUID"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetResampledData: &mockGetResampledData{
				res: nil,
				err: errors.New("error"),
			},
			wantStatusCode: 500,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)
			translator := infrastructure.NewTranslator()
			mockCustomValidator := new(mocks.CustomValidator)
			validator := infrastructure.NewValidator(translator, mockCustomValidator)
			e := pkg.NewRouter()
			e.Validator = validator

			req := httptest.NewRequest(http.MethodPost, "/iot/sensors/resampledData/:nodeId", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/iot/sensors/resampledData/:nodeId")
			c.SetParamNames("nodeId")
			c.SetParamValues("UUID")

			mockSensorService := new(mocks.SensorService)

			if tt.mockGetResampledData != nil {
				mockSensorService.On("GetResampledData", mock.Anything, mock.Anything).Return(tt.mockGetResampledData.res, tt.mockGetResampledData.err)
			}
			if tt.mockCustomValidator != nil {
				mockCustomValidator.On("ExistValidator", mock.Anything).Return(tt.mockCustomValidator.res)
			}

			handler := &Controller{
				Logger:        logger,
				Translator:    translator,
				SensorService: mockSensorService,
			}

			err := handler.HandleGetResampledData(c)

			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tt.wantStatusCode, he.Code)
				} else {
					assert.Equal(t, tt.wantErr, err != nil)
				}
			} else {
				assert.Equal(t, tt.wantStatusCode, rec.Code)
			}
		})
	}
}
