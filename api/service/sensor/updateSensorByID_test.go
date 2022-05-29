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

func Test_service_UpdateSensorByID(t *testing.T) {

	type args struct {
		ctx     context.Context
		payload *model.UpdateSensorRequest
	}
	type mockGetSensorByID struct {
		res *database.Sensor
		err error
	}
	type mockUpdateSensorByID struct {
		err error
	}
	tests := []struct {
		name                 string
		args                 args
		mockGetSensorByID    *mockGetSensorByID
		mockUpdateSensorByID *mockUpdateSensorByID
		wantErr              bool
	}{
		{
			name: "WHEN update sensor data by ID THEN system return success",
			args: args{
				ctx: context.TODO(),
				payload: &model.UpdateSensorRequest{
					ID:          "UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: &database.Sensor{
					ID:     "UUID",
					NodeID: "NODE-UUID",
				},
				err: nil,
			},
			mockUpdateSensorByID: &mockUpdateSensorByID{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "WHEN update sensor data by ID AND sensor not found THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.UpdateSensorRequest{
					ID:          "UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "WHEN update sensor data by ID AND got error WHEN find sensor by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.UpdateSensorRequest{
					ID:          "UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: nil,
				err: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "WHEN update sensor data by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.UpdateSensorRequest{
					ID:          "UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: &database.Sensor{
					ID:     "UUID",
					NodeID: "NODE-UUID",
				},
				err: nil,
			},
			mockUpdateSensorByID: &mockUpdateSensorByID{
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
			if tt.mockUpdateSensorByID != nil {
				mockSensorRepo.On("UpdateSensorByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockUpdateSensorByID.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
				},
			}
			if err := svc.UpdateSensorByID(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateSensorByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
