package model

type Item struct {
	ID     int    `json:"id" gorm:"type:int;autoIncrement"`
	Name   string `json:"name" gorm:"type:varchar(255);not null"`
	UserID int    `json:"user_id" gorm:"type:int;not null"`
	User   User   `json:"-" gorm:"constraint:onUpdate:CASCADE,OnDelete:CASCADE;"`
}
