package model

// TODO: MOVE TO BUILDING BLOCK
import (
	"github.com/araddon/dateparse"
	"strings"
	"time"
)

type DateColumn struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type DateTime time.Time

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := dateparse.ParseAny(value)

	if err != nil {
		return err
	}
	*dt = DateTime(t)
	return nil
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(dt).Format("2006-01-02 15:04:05") + `"`), nil
}
