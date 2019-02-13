package entity

import (
	"database/sql"
)


type Order_activity struct {
	Id string
	Discount sql.NullString
	Full_minus sql.NullString
	Full_give sql.NullString
	Work string
	Created_time string
	End_time string
}