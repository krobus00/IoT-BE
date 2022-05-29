package node

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

func Test_service_GetNodeInfo(t *testing.T) {
	type args struct {
		ctx    context.Context
		nodeID string
	}

	type mockFindNodeByID struct {
		res *database.Node
		err error
	}

	type mockGetLastReportByNodeID struct {
		res *database.Sensor
		err error
	}

	tests := []struct {
		name                      string
		args                      args
		mockFindNodeByID          *mockFindNodeByID
		mockGetLastReportByNodeID *mockGetLastReportByNodeID
		want                      *model.GetNodeInfoResponse
		wantErr                   bool
	}{
		{
			name: "WHEN get current node info THEN system return node info",
			args: args{
				ctx:    context.TODO(),
				nodeID: "UUID",
			},
			mockFindNodeByID: &mockFindNodeByID{
				res: &database.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10,
					Latitude:   17,
					ModelURL:   "https://bucket-name",
					DateColumn: database.DateColumn{},
				},
				err: nil,
			},
			mockGetLastReportByNodeID: &mockGetLastReportByNodeID{
				res: &database.Sensor{
					ID:          "UUID",
					NodeID:      "UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
					DateColumn:  database.DateColumn{},
				},
				err: nil,
			},
			want: &model.GetNodeInfoResponse{
				NodeResponse: model.NodeResponse{
					ID:         "UUID",
					City:       "city",
					Longitude:  10,
					Latitude:   17,
					DateColumn: model.DateColumn{},
				},
				LastReport: &model.SensorResponse{
					ID:          "UUID",
					NodeID:      "UUID",
					Humidity:    10,
					Temperature: 17,
					HeatIndex:   30,
					DateColumn:  model.DateColumn{},
				},
			},
			wantErr: false,
		},
		{
			name: "WHEN get current node info AND node not found THEN system return error",
			args: args{
				ctx:    context.TODO(),
				nodeID: "UUID",
			},
			mockFindNodeByID: &mockFindNodeByID{
				res: nil,
				err: nil,
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "WHEN get current node info AND last report not found THEN system return node info",
			args: args{
				ctx:    context.TODO(),
				nodeID: "UUID",
			},
			mockFindNodeByID: &mockFindNodeByID{
				res: &database.Node{
					ID:         "UUID",
					City:       "city",
					Longitude:  10,
					Latitude:   17,
					ModelURL:   "https://bucket-name",
					DateColumn: database.DateColumn{},
				},
				err: nil,
			},
			mockGetLastReportByNodeID: &mockGetLastReportByNodeID{
				res: nil,
				err: nil,
			},
			want: &model.GetNodeInfoResponse{
				NodeResponse: model.NodeResponse{
					ID:         "UUID",
					City:       "city",
					Longitude:  10,
					Latitude:   17,
					DateColumn: model.DateColumn{},
				},
				LastReport: &model.SensorResponse{
					NodeID: "UUID",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx = context.WithValue(tt.args.ctx, "nodeId", tt.args.nodeID)
			sqlxDB, _ := sqlx.Open("test", "test")
			config := infrastructure.Env{}
			logger := infrastructure.NewLogger(config)

			mockNodeRepo := new(mocks.NodeRepository)
			mockSensorRepo := new(mocks.SensorRepository)

			if tt.mockFindNodeByID != nil {
				mockNodeRepo.On("FindNodeByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockFindNodeByID.res, tt.mockFindNodeByID.err)
			}
			if tt.mockGetLastReportByNodeID != nil {
				mockSensorRepo.On("GetLastReportByNodeID", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetLastReportByNodeID.res, tt.mockGetLastReportByNodeID.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					NodeRepository:   mockNodeRepo,
					SensorRepository: mockSensorRepo,
				},
			}
			got, err := svc.GetNodeInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetNodeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetNodeInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
