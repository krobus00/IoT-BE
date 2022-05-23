package sensor

import (
	"testing"

	"github.com/krobus00/iot-be/infrastructure"
)

func Test_repository_GetTableName(t *testing.T) {
	type fields struct {
		logger infrastructure.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "sensors",
			fields: fields{},
			want:   "sensors",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				logger: tt.fields.logger,
			}
			if got := r.GetTableName(); got != tt.want {
				t.Errorf("repository.GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
