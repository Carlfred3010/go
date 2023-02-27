package models

import "time"

type Stock struct {
	Symbol    string `json:"symbol" gorm:"primary_key"`
	CreatedAt time.Time
	Name      string  `json:"name" gorm:"type:varchar(255)"`
	Price     float64 `json:"price" gorm:"type:float(32)"`
}
