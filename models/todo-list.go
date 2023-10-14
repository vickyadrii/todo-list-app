package models

import "time"

type TodoList struct {
	Id        uint   `gorm:"primaryKey" json:"id" form:"id"`
	List      string `gorm:"not null" json:"list"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
