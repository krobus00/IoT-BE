package database

type Node struct {
	ID        string  `db:"id"`
	City      string  `db:"city"`
	Longitude float64 `db:"longitude"`
	Latitude  float64 `db:"latitude"`
	ModelURL  string  `db:"model_url"`
	DateColumn
}
