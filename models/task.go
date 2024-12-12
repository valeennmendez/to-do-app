package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID     uint
	UserID int    `json:"userid" gorm:"foreignkey:UserID""`
	Task   string `json:"task" validate:"requiered"`
	User   User   `json: "user" gorm:"association_foreignkey:UserID"`
}
