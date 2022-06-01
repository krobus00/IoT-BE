package sensor

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	"github.com/stretchr/testify/mock"
)

func Test_service_GetAllSensor(t *testing.T) {

	type args struct {
		ctx     context.Context
		payload *kro_model.PaginationRequest
	}
	type mockGetAllSensor struct {
		res []*database.Sensor
		err error
	}
	type mockCountSensors struct {
		res int64
		err error
	}
	tests := []struct {
		name             string
		args             args
		mockGetAllSensor *mockGetAllSensor
		mockCountSensors *mockCountSensors
		want             *kro_model.PaginationResponse
		wantErr          bool
	}{
		{
			name: "WHEN get all sensors THEN system return paginated sensors",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllSensor: &mockGetAllSensor{
				res: []*database.Sensor{
					{
						ID:          "UUID",
						NodeID:      "NODE-UUID",
						Humidity:    70,
						Temperature: 10,
						HeatIndex:   30,
						DateColumn:  database.DateColumn{},
					},
				},
				err: nil,
			},
			mockCountSensors: &mockCountSensors{
				res: 2,
				err: nil,
			},
			want: &kro_model.PaginationResponse{
				CurrentPage:  1,
				ItemsPerPage: 1,
				Count:        2,
				TotalPage:    2,
				Items: []*model.SensorResponse{
					{
						ID:          "UUID",
						NodeID:      "NODE-UUID",
						Humidity:    70,
						Temperature: 10,
						HeatIndex:   30,
						DateColumn:  model.DateColumn{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "WHEN get all sensors AND got error WHEN fetch sensors THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllSensor: &mockGetAllSensor{
				res: nil,
				err: errors.New("error"),
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "WHEN get all sensors AND got error WHEN count sensors THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllSensor: &mockGetAllSensor{
				res: []*database.Sensor{
					{
						ID:          "UUID",
						NodeID:      "NODE-UUID",
						Humidity:    70,
						Temperature: 10,
						HeatIndex:   30,
						DateColumn:  database.DateColumn{},
					},
				},
				err: nil,
			},
			mockCountSensors: &mockCountSensors{
				res: 0,
				err: errors.New("error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlxDB, _ := sqlx.Open("test", "test")
			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)

			mockSensorRepo := new(mocks.SensorRepository)

			if tt.mockGetAllSensor != nil {
				mockSensorRepo.On("GetAllSensor", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetAllSensor.res, tt.mockGetAllSensor.err)
			}

			if tt.mockCountSensors != nil {
				mockSensorRepo.On("CountSensors", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockCountSensors.res, tt.mockCountSensors.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
				},
			}
			got, err := svc.GetAllSensor(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAllSensor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAllSensor() = %v, want %v", got, tt.want)
			}
		})
	}
}
