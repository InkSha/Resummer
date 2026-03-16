package model

import "gorm.io/gorm"

type UserStatus int8

const (
	UserDeactivate UserStatus = 0
	UserActivate   UserStatus = 1
)

type User struct {
	gorm.Model

	Account  string     `gorm:"uniqueIndex;not null"`
	Password string     `gorm:"not null;size:60"`
	Status   UserStatus `gorm:"default:1;not null;"`
}
