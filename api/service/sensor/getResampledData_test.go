package sensor

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/api/requester"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	"github.com/stretchr/testify/mock"
)

func Test_service_GetResampledData(t *testing.T) {
	var t1, t2, t3 float64
	t1, t2, t3 = 10, 17, 30
	tn := time.Now().Unix()
	type args struct {
		ctx     context.Context
		payload *model.GetProcessedDataRequest
	}
	type mockGetSensorByRange struct {
		res []*database.Sensor
		err error
	}
	type mockCallResamplingData struct {
		res []*model.GetSampledData
		err error
	}
	tests := []struct {
		name                   string
		args                   args
		mockGetSensorByRange   *mockGetSensorByRange
		mockCallResamplingData *mockCallResamplingData
		want                   *model.GetProcessedDataResponse
		wantErr                bool
	}{
		{
			name: "WHEN get resampled data THEN system return success",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetProcessedDataRequest{
					NodeID:    "NODE-UUID",
					StartDate: tn,
					EndDate:   tn,
				},
			},
			mockGetSensorByRange: &mockGetSensorByRange{
				res: []*database.Sensor{
					{
						ID:     "UUID",
						NodeID: "NODE-UUID",
					},
				},
				err: nil,
			},
			mockCallResamplingData: &mockCallResamplingData{
				res: []*model.GetSampledData{
					{
						Humidity:    t1,
						Temperature: t2,
						HeatIndex:   t3,
						CreatedAt:   tn,
					},
				},
			},
			want: &model.GetProcessedDataResponse{
				Humidity:    []*float64{&t1},
				Temperature: []*float64{&t2},
				HeatIndex:   []*float64{&t3},
				DateTime:    []*int64{&tn},
			},
			wantErr: false,
		},
		{
			name: "WHEN get resampled data AND got error WHEN get sensor data THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetProcessedDataRequest{
					NodeID:    "NODE-UUID",
					StartDate: tn,
					EndDate:   tn,
				},
			},
			mockGetSensorByRange: &mockGetSensorByRange{
				res: nil,
				err: errors.New("error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WHEN get resampled data  AND got error WHEN send resampling request THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetProcessedDataRequest{
					NodeID:    "NODE-UUID",
					StartDate: tn,
					EndDate:   tn,
				},
			},
			mockGetSensorByRange: &mockGetSensorByRange{
				res: []*database.Sensor{
					{
						ID:     "UUID",
						NodeID: "NODE-UUID",
					},
				},
				err: nil,
			},
			mockCallResamplingData: &mockCallResamplingData{
				res: nil,
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
			mockDataRequester := new(mocks.DataRequester)

			if tt.mockGetSensorByRange != nil {
				mockSensorRepo.On("GetSensorByRange", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetSensorByRange.res, tt.mockGetSensorByRange.err)
			}

			if tt.mockCallResamplingData != nil {
				mockDataRequester.On("CallResamplingData", mock.Anything, mock.Anything).Return(tt.mockCallResamplingData.res, tt.mockCallResamplingData.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				requester: requester.Requester{
					DataRequester: mockDataRequester,
				},
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
				},
			}
			got, err := svc.GetResampledData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetResampledData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetResampledData() = %v, want %v", got, tt.want)
			}
		})
	}
}
