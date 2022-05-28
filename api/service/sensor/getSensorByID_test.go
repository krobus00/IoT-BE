package sensor

import (
	"context"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	"github.com/stretchr/testify/mock"
)

func Test_service_GetSensorByID(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *model.ShowSensorRequest
	}
	type mockGetSensorByID struct {
		res *database.Sensor
		err error
	}
	tests := []struct {
		name              string
		args              args
		mockGetSensorByID *mockGetSensorByID
		want              *model.SensorResponse
		wantErr           bool
	}{
		{
			name: "WHEN get sensor data by id THEN system return sensor data",
			args: args{
				ctx: context.TODO(),
				payload: &model.ShowSensorRequest{
					ID: "UUID",
				},
			},
			mockGetSensorByID: &mockGetSensorByID{
				res: &database.Sensor{
					ID:          "UUID",
					NodeID:      "NODE-UUID",
					Humidity:    70,
					Temperature: 10,
					HeatIndex:   30,
					DateColumn:  database.DateColumn{},
				},
				err: nil,
			},
			want: &model.SensorResponse{
				ID:          "UUID",
				NodeID:      "NODE-UUID",
				Humidity:    70,
				Temperature: 10,
				HeatIndex:   30,
				DateColumn:  model.DateColumn{},
			},
			wantErr: false,
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

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
				},
			}
			got, err := svc.GetSensorByID(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetSensorByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetSensorByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
