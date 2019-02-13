package entity

import (
	"database/sql"
)


type Dishes struct {
	Did string
	Name string
	Price string
	// NULL 时是 false
	Img sql.NullString	// default 为 NULL 要用这个
	Type_id string
	Type_name string
	Status string
}