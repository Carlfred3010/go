package models

import ("time"
	
)
type User struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time

	FirstName string `json:"first_name" gorm:"type:varchar(255)"`
	Lastname  string `json:"last_name" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255)"`
}
