package sensor

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/stretchr/testify/mock"
)

func Test_service_StoreSensor(t *testing.T) {

	type args struct {
		ctx     context.Context
		payload *model.CreateSensorRequest
		nodeID  string
	}
	type mockStore struct {
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockStore *mockStore
		wantErr   bool
	}{
		{
			name: "WHEN store new sensor data THEN system store new sensor into DB",
			args: args{
				ctx: context.TODO(),
				payload: &model.CreateSensorRequest{
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
				},
				nodeID: "UUID",
			},
			mockStore: &mockStore{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN store new sensor data THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.CreateSensorRequest{
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
				},
				nodeID: "UUID",
			},
			mockStore: &mockStore{
				err: errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = context.WithValue(tt.args.ctx, "nodeId", tt.args.nodeID)
			sqlxDB, _ := sqlx.Open("test", "test")
			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)

			mockSensorRepo := new(mocks.SensorRepository)

			if tt.mockStore != nil {
				mockSensorRepo.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockStore.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
				},
			}
			if err := svc.StoreSensor(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("service.StoreSensor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
