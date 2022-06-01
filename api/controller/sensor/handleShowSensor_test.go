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

func TestController_HandleShowSensor(t *testing.T) {
	type mockGetSensorByID struct {
		res *model.SensorResponse
		err error
	}
	type mockCustomValidator struct {
		res bool
	}
	tests := []struct {
		name                string
		payload             string
		mockCustomValidator *mockCustomValidator
		mockGetSensorByID   *mockGetSensorByID
		wantStatusCode      int
		wantErr             bool
	}{
		{
			name: "WHEN get node info by id THEN system return node info",
			payload: `{
				"id": "55a21501-864f-4c19-adf5-d7b9db1d3aa3"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: &model.SensorResponse{},
			},
			wantStatusCode: 200,
			wantErr:        false,
		},
		{
			name:    "WHEN get node info by id AND payload malformated THEN system return error",
			payload: `{`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name: "WHEN get node info by id AND payload invalid THEN system return error",
			payload: `{
				"id": "55a21501-864f-4c19-adf5-d7b9db1d3aa3"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: false,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name: "WHEN get node info by id THEN system return error",
			payload: `{
				"id": "55a21501-864f-4c19-adf5-d7b9db1d3aa3"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetSensorByID: &mockGetSensorByID{
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
			req := httptest.NewRequest(http.MethodPost, "/iot/sensor/:id", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// set echo param
			c.SetPath("/iot/sensor/:id")
			c.SetParamNames("id")
			c.SetParamValues("UUID")

			mockSensorService := new(mocks.SensorService)

			if tt.mockGetSensorByID != nil {
				mockSensorService.On("GetSensorByID", mock.Anything, mock.Anything).Return(tt.mockGetSensorByID.res, tt.mockGetSensorByID.err)
			}
			if tt.mockCustomValidator != nil {
				mockCustomValidator.On("ExistValidator", mock.Anything).Return(tt.mockCustomValidator.res)
			}

			handler := &Controller{
				Logger:        logger,
				Translator:    translator,
				SensorService: mockSensorService,
			}

			err := handler.HandleShowSensor(c)

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
