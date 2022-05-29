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

func Test_service_GetForecastData(t *testing.T) {
	var t1, t2 float64
	t1, t2 = 10, 17
	tn := time.Now().Unix()

	type args struct {
		ctx     context.Context
		payload *model.GetForecastDataRequest
	}
	type mockGetSensorByRange struct {
		res []*database.Sensor
		err error
	}
	type mockFindNodeByID struct {
		res *database.Node
		err error
	}
	type mockCallForecastData struct {
		res []*model.GetForecastData
		err error
	}
	tests := []struct {
		name                 string
		args                 args
		mockGetSensorByRange *mockGetSensorByRange
		mockFindNodeByID     *mockFindNodeByID
		mockCallForecastData *mockCallForecastData
		want                 *model.GetForecastDataResponse
		wantErr              bool
	}{
		{
			name: "WHEN get forecasting data THEN system return 24h Forecast data",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetForecastDataRequest{
					NodeID: "NODE-UUID",
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
			mockFindNodeByID: &mockFindNodeByID{
				res: &database.Node{
					ID:         "NODE-UUID",
					City:       "city",
					Longitude:  10,
					Latitude:   17,
					ModelURL:   "https://bucket-name",
					DateColumn: database.DateColumn{},
				},
			},
			mockCallForecastData: &mockCallForecastData{
				res: []*model.GetForecastData{
					{
						Temperature: t1,
						CreatedAt:   tn,
					},
					{
						Temperature: t2,
						CreatedAt:   tn,
					},
				},
			},
			want: &model.GetForecastDataResponse{
				Temperature: []*float64{&t1, &t2},
				DateTime:    []*int64{&tn, &tn},
			},
			wantErr: false,
		},
		{
			name: "WHEN get forecasting data AND got error when fetch sensor data THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetForecastDataRequest{
					NodeID: "NODE-UUID",
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
			name: "WHEN get forecasting data AND sensor data is empty THEN system return nil",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetForecastDataRequest{
					NodeID: "NODE-UUID",
				},
			},
			mockGetSensorByRange: &mockGetSensorByRange{
				res: nil,
				err: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "WHEN get forecasting data AND got error WHEN find node by ID THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetForecastDataRequest{
					NodeID: "NODE-UUID",
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
			mockFindNodeByID: &mockFindNodeByID{
				res: nil,
				err: errors.New("error"),
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "WHEN get forecasting data AND got error WHEN send forecast request THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &model.GetForecastDataRequest{
					NodeID: "NODE-UUID",
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
			mockFindNodeByID: &mockFindNodeByID{
				res: &database.Node{
					ID:         "NODE-UUID",
					City:       "city",
					Longitude:  10,
					Latitude:   17,
					ModelURL:   "https://bucket-name",
					DateColumn: database.DateColumn{},
				},
			},
			mockCallForecastData: &mockCallForecastData{
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
			mockNodeRepo := new(mocks.NodeRepository)
			mockDataRequester := new(mocks.DataRequester)

			if tt.mockGetSensorByRange != nil {
				mockSensorRepo.On("GetSensorByRange", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetSensorByRange.res, tt.mockGetSensorByRange.err)
			}

			if tt.mockFindNodeByID != nil {
				mockNodeRepo.On("FindNodeByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFindNodeByID.res, tt.mockFindNodeByID.err)
			}

			if tt.mockCallForecastData != nil {
				mockDataRequester.On("CallForecastData", mock.Anything, mock.Anything).Return(tt.mockCallForecastData.res, tt.mockCallForecastData.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				requester: requester.Requester{
					DataRequester: mockDataRequester,
				},
				repository: repository.Repository{
					SensorRepository: mockSensorRepo,
					NodeRepository:   mockNodeRepo,
				},
			}
			got, err := svc.GetForecastData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetForecastData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetForecastData() = %v, want %v", got, tt.want)
			}
		})
	}
}
