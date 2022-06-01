package node

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/krobot-building-block/pkg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_HandleGetNode(t *testing.T) {
	type mockGetNodeInfo struct {
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
		mockGetNodeInfo     *mockGetNodeInfo
		wantStatusCode      int
		wantErr             bool
	}{
		{
			name: "WHEN get node info THEN system return success",
			payload: `{
				"id": "55a21501-864f-4c19-adf5-d7b9db1d3aa3"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetNodeInfo: &mockGetNodeInfo{
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
			name: "WHEN get node info THEN system return error",
			payload: `{
				"id": "55a21501-864f-4c19-adf5-d7b9db1d3aa3"
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockGetNodeInfo: &mockGetNodeInfo{
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
			req := httptest.NewRequest(http.MethodGet, "/iot/nodes/me", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// set echo context
			c.Set("nodeId", "UUID")

			mockNodeService := new(mocks.NodeService)

			if tt.mockGetNodeInfo != nil {
				mockNodeService.On("GetNodeInfo", mock.Anything).Return(tt.mockGetNodeInfo.res, tt.mockGetNodeInfo.err)
			}
			if tt.mockCustomValidator != nil {
				mockCustomValidator.On("ExistValidator", mock.Anything).Return(tt.mockCustomValidator.res)
			}

			handler := &Controller{
				Logger:      logger,
				Translator:  translator,
				NodeService: mockNodeService,
			}

			err := handler.HandleGetNode(c)

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
