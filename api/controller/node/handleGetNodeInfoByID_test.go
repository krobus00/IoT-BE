package node

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

func TestController_HandleGetNodeInfoByID(t *testing.T) {
	type mockGetNodeInfoByID struct {
		res *model.GetNodeInfoResponse
		err error
	}
	type mockCustomValidator struct {
		res bool
	}
	tests := []struct {
		name                string
		payload             string
		mockCustomValidator *mockCustomValidator
		mockGetNodeInfoByID *mockGetNodeInfoByID
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
			mockGetNodeInfoByID: &mockGetNodeInfoByID{
				res: &model.GetNodeInfoResponse{
					NodeResponse: model.NodeResponse{
						ID: "UUID",
					},
					LastReport: &model.SensorResponse{
						ID: "UUID",
					},
				},
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
			mockGetNodeInfoByID: &mockGetNodeInfoByID{
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
			req := httptest.NewRequest(http.MethodPost, "/iot/nodes/:id", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// set echo param
			c.SetPath("/iot/nodes/:id")
			c.SetParamNames("id")
			c.SetParamValues("UUID")

			mockNodeService := new(mocks.NodeService)

			if tt.mockGetNodeInfoByID != nil {
				mockNodeService.On("GetNodeInfoByID", mock.Anything, mock.Anything).Return(tt.mockGetNodeInfoByID.res, tt.mockGetNodeInfoByID.err)
			}
			if tt.mockCustomValidator != nil {
				mockCustomValidator.On("ExistValidator", mock.Anything).Return(tt.mockCustomValidator.res)
			}

			handler := &Controller{
				Logger:      logger,
				Translator:  translator,
				NodeService: mockNodeService,
			}

			err := handler.HandleGetNodeInfoByID(c)

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
