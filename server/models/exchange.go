package models

import "time"

type Exchange struct {
	ID        uint      `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	Type      string    `json:"type" gorm:"column:type"`
	Bid       float64   `json:"bid" gorm:"column:bid"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
