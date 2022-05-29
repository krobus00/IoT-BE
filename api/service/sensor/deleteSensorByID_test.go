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
	"github.com/krobus00/iot-be/model/database"
	"github.com/stretchr/testify/mock"
)

func Test_service_DeleteSensorByID(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *model.DeleteSensorRequest
	}
	type mockGetSensorByID struct {
		res *database.Sensor
		err error
	}
	type mockDeleteSensorByID struct {
		err error
	}
	tests := []struct {
		name                 string
		args                 args
		mockGetSensorByID    *mockGetSensorByID
		mockDeleteSensorByID *mockDeleteSensorByID
		wantErr              bool
	}{
		{
			name: "WHEN delete sensor data by ID THEN system return success",
			args: args{
				ctx: context.TODO(),
				payload: &model.DeleteSensorRequest{
					ID: "UUID",
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: &database.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
					DateColumn:  database.DateColumn{},
				},
				err: nil,
			},
			mockDeleteSensorByID: &mockDeleteSensorByID{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN delete sensor data by ID AND sensor not found THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.DeleteSensorRequest{
					ID: "UUID",
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: nil,
				err: nil,
			},
			mockDeleteSensorByID: nil,
			wantErr:              true,
		},
		{
			name: "WHEN delete sensor data by ID AND got error WHEN get sensor data by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.DeleteSensorRequest{
					ID: "UUID",
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: nil,
				err: errors.New("error"),
			},
			mockDeleteSensorByID: nil,
			wantErr:              true,
		},
		{
			name: "WHEN delete sensor data by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.DeleteSensorRequest{
					ID: "UUID",
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: &database.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
					DateColumn:  database.DateColumn{},
				},
				err: nil,
			},
			mockDeleteSensorByID: &mockDeleteSensorByID{
				err: errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")
			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)

			mockSensorRepo := new(mocks.SensorRepository)

			if tt.mockGetSensorByID != nil {
				mockSensorRepo.On("GetSensorByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetSensorByID.res, tt.mockGetSensorByID.err)
			}
			if tt.mockDeleteSensorByID != nil {
				mockSensorRepo.On("DeleteSensorByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockDeleteSensorByID.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
				},
			}
			if err := svc.DeleteSensorByID(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteSensorByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
