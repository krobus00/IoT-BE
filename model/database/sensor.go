package database

type Sensor struct {
	ID          string  `db:"id"`
	NodeID      string  `db:"node_id"`
	Humidity    float64 `db:"humidity"`
	Temperature float64 `db:"temperature"`
	HeatIndex   float64 `db:"heat_index"`
	DateColumn
}
