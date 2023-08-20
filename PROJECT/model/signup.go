package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}
type Signup struct {
	Name             string `json:"name" `
	Email            string `json:"email" binding:"required,email" gorm:"not null;unique"`
	PhoneNumber      string `json:"phonenumber" valid:"required,minlength(10),maxlength(10)"`
	Lengthofpassword int    `json:"lengthofpassword" binding:"required"`
}
