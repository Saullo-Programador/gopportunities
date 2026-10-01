package schemas

import ("gorm.io/gorm")

type Opening struct {
	gorm.Model
	Role string
	Title string
	Description string
	Company string
	Location string
	Remote bool
	Link string
	Salary string
}