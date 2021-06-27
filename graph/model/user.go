package model

type User struct {
	ID   int    `json:"id" gorm:"type:int;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
}


