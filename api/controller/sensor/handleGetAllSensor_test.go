package sensor

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	kro_model "github.com/krobus00/krobot-building-block/model"
	"github.com/krobus00/krobot-building-block/pkg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_HandleGetAllSensor(t *testing.T) {
	type mockGetAllSensor struct {
		res *kro_model.PaginationResponse
		err error
	}
	type mockCustomValidator struct {
		res bool
	}
	tests := []struct {
		name                string
		pageSize            string
		mockCustomValidator *mockCustomValidator
		mockGetAllSensor    *mockGetAllSensor
		wantStatusCode      int
		wantErr             bool
	}{
		{
			name:     "WHEN get all sensor THEN system return paginate node",
			pageSize: "1",
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetAllSensor: &mockGetAllSensor{
				res: &kro_model.PaginationResponse{
					CurrentPage:  1,
					ItemsPerPage: 1,
					Count:        1,
					TotalPage:    1,
				},
			},
			wantStatusCode: 200,
			wantErr:        false,
		},
		{
			name:     "WHEN get all sensor AND payload malformated THEN system return error",
			pageSize: "aaa",
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name:     "WHEN get all sensor THEN system return error",
			pageSize: "1",
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetAllSensor: &mockGetAllSensor{
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

			// set query params
			q := make(url.Values)
			q.Set("pageSize", tt.pageSize)

			req := httptest.NewRequest(http.MethodGet, "/iot/sensor?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockSensorService := new(mocks.SensorService)

			if tt.mockGetAllSensor != nil {
				mockSensorService.On("GetAllSensor", mock.Anything, mock.Anything).Return(tt.mockGetAllSensor.res, tt.mockGetAllSensor.err)
			}
			if tt.mockCustomValidator != nil {
				mockCustomValidator.On("ExistValidator", mock.Anything).Return(tt.mockCustomValidator.res)
			}

			handler := &Controller{
				Logger:        logger,
				Translator:    translator,
				SensorService: mockSensorService,
			}

			err := handler.HandleGetAllSensor(c)

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
