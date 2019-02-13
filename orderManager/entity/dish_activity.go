package entity

import (
	"database/sql"
)


type Dish_activity struct {
	Id string
	Did string
	Name string
	Work string
	Minus_price sql.NullString
	Discount sql.NullString
	Created_time string
	End_time string
}