package node

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/krobot-building-block/pkg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_HandleRegister(t *testing.T) {
	uuid := "UUID UUIOD"
	type mockRegister struct {
		res *string
		err error
	}
	type mockCustomValidator struct {
		res bool
	}
	tests := []struct {
		name                string
		payload             string
		mockCustomValidator *mockCustomValidator
		mockRegister        *mockRegister
		wantStatusCode      int
		wantErr             bool
	}{
		{
			name: "WHEN register new node THEN system return success",
			payload: `{
				"city": "test node 17",
				"longitude": 17.10,
				"latitude": 10.17
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockRegister: &mockRegister{
				res: &uuid,
				err: nil,
			},
			wantStatusCode: 200,
			wantErr:        false,
		},
		{
			name:    "WHEN register new node AND payload malformated THEN system return error",
			payload: `{`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name: "WHEN register new node AND payload invalid THEN system return error",
			payload: `{
				"city": "test node 17",
				"longitude": 17.10,
				"latitude": 10.17
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: false,
			},
			wantStatusCode: 400,
			wantErr:        true,
		},
		{
			name: "WHEN register new node THEN system return error",
			payload: `{
				"city": "test node 17",
				"longitude": 17.10,
				"latitude": 10.17
			}`,
			mockCustomValidator: &mockCustomValidator{
				res: true,
			},
			mockRegister: &mockRegister{
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

			req := httptest.NewRequest(http.MethodPost, "/_internal/iot/nodes/register", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockNodeService := new(mocks.NodeService)

			if tt.mockRegister != nil {
				mockNodeService.On("Register", mock.Anything, mock.Anything).Return(tt.mockRegister.res, tt.mockRegister.err)
			}
			if tt.mockCustomValidator != nil {
				mockCustomValidator.On("UniqueValidator", mock.Anything).Return(tt.mockCustomValidator.res)
			}

			handler := &Controller{
				Logger:      logger,
				Translator:  translator,
				NodeService: mockNodeService,
			}

			err := handler.HandleRegister(c)

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
