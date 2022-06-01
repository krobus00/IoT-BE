package sensor

import (
	"testing"

	"github.com/krobus00/iot-be/infrastructure"
)

func Test_repository_GetTableName(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "sensors",
			want: "sensors",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(infrastructure.Infrastructure{})
			if got := r.GetTableName(); got != tt.want {
				t.Errorf("repository.GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
