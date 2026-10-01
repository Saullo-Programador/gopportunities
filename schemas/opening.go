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

type OpeningResponse struct {
	ID uint `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	Role string `json:"role"`
	Title string `json:"title"`
	Description string `json:"description"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Salary string `json:"salary"`
}