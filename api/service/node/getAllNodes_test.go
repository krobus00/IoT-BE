package node

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/iot-be/api/repository"
	"github.com/krobus00/iot-be/infrastructure"
	"github.com/krobus00/iot-be/mocks"
	"github.com/krobus00/iot-be/model"
	"github.com/krobus00/iot-be/model/database"
	kro_model "github.com/krobus00/krobot-building-block/model"
	"github.com/stretchr/testify/mock"
)

func Test_service_GetAllNodes(t *testing.T) {
	tn := time.Now().Unix()
	type args struct {
		ctx     context.Context
		payload *kro_model.PaginationRequest
	}

	type mockGetAllNodes struct {
		res []*database.Node
		err error
	}
	type mockCountNodes struct {
		res int64
		err error
	}
	tests := []struct {
		name            string
		args            args
		mockGetAllNodes *mockGetAllNodes
		mockCountNodes  *mockCountNodes
		want            *kro_model.PaginationResponse
		wantErr         bool
	}{
		{
			name: "WHEN get all nodes THEN system return paginated nodes",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllNodes: &mockGetAllNodes{
				res: []*database.Node{
					{
						ID:        "UUID",
						City:      "city",
						Longitude: 10.10,
						Latitude:  17.17,
						ModelURL:  "https://bucket_url",
						DateColumn: database.DateColumn{
							CreatedAt: tn,
							UpdatedAt: tn,
							DeletedAt: nil,
						},
					},
				},
				err: nil,
			},
			mockCountNodes: &mockCountNodes{
				res: 2,
				err: nil,
			},
			want: &kro_model.PaginationResponse{
				CurrentPage:  1,
				ItemsPerPage: 1,
				Count:        2,
				TotalPage:    2,
				Items: []*model.NodeResponse{
					{
						ID:        "UUID",
						City:      "city",
						Longitude: 10.10,
						Latitude:  17.17,
						DateColumn: model.DateColumn{
							CreatedAt: tn,
							UpdatedAt: tn,
							DeletedAt: nil,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "WHEN get all nodes AND got error WHEN fetch nodes THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllNodes: &mockGetAllNodes{
				res: nil,
				err: errors.New("error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WHEN get all nodes AND nodes not found THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllNodes: &mockGetAllNodes{
				res: nil,
				err: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "WHEN get all nodes AND got error WHEN count all nodes THEN system return error",
			args: args{
				ctx: context.TODO(),
				payload: &kro_model.PaginationRequest{
					PageSize: 1,
					Page:     1,
				},
			},
			mockGetAllNodes: &mockGetAllNodes{
				res: []*database.Node{
					{
						ID:        "UUID",
						City:      "city",
						Longitude: 10.10,
						Latitude:  17.17,
						ModelURL:  "https://bucket_url",
						DateColumn: database.DateColumn{
							CreatedAt: tn,
							UpdatedAt: tn,
							DeletedAt: nil,
						},
					},
				},
				err: nil,
			},
			mockCountNodes: &mockCountNodes{
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

			mockNodeRepo := new(mocks.NodeRepository)

			if tt.mockGetAllNodes != nil {
				mockNodeRepo.On("GetAllNodes", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockGetAllNodes.res, tt.mockGetAllNodes.err)
			}

			if tt.mockCountNodes != nil {
				mockNodeRepo.On("CountNodes", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockCountNodes.res, tt.mockCountNodes.err)
			}

			svc := &service{
				logger: logger,
				db:     sqlxDB,
				repository: repository.Repository{
					NodeRepository: mockNodeRepo,
				},
			}
			got, err := svc.GetAllNodes(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAllNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAllNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
