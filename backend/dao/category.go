package dao

import "database/sql"

type Category struct {
	ID    int32          `json:"id"`
	Name  string         `jsong:"name"`
	Image sql.NullString `json:"image"`
}
