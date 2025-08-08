package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;size:50" json:"email"`
	Password  string `gorm:"size:100" json:"-"`
	FirstName string `gorm:"size:50" json:"first_name"`
	LastName  string `gorm:"size:50" json:"last_name"`
}

type RegisterReq struct {
	Email     string `json:"email" validate:"required,email,max=50"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	FirstName string `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" validate:"required,min=1,max=50"`
}
