package dao

import "database/sql"

type Furniture struct {
	ID       int32          `json:"id"`
	Name     string         `json:"name"`
	Category string         `json:"category"`
	Price    float64        `json:"price"`
	Image    sql.NullString `json:"image"`
}
