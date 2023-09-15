package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	Id       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email" gorm:"index"`
	Password string `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Perform some actions before creating a use
	u.Role = "super admin"
	fmt.Println("Preparing to create user:", u)
	return nil
}
