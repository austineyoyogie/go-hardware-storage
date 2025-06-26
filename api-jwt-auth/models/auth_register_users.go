package models

import (
	"errors"
	"regexp"
)

var (
	ErrEmptyFields = errors.New("User field can't be empty") 
	ErrInvalidEmail = errors.New("User email is not a valid email address")
	ErrInvalidTelephone = errors.New("User telephone is not a valid number")
	ErrInvalidPassword = errors.New("User password is the wrong length (should be 8~20 characters)")
)

type User struct {
	Model
	FirstName    string `gorm:"size:45;not null;" json:"first_name"`
	LastName     string `gorm:"size:45;not null;" json:"last_name"`
	Email 	     string `gorm:"size:45;not null;unique" json:"email"`
	Password	 string `gorm:"size:250;not null;" json:"password"`
	Telephone	 string `gorm:"size:45;not null;unique" json:"telephone"`
	Token        string `gorm:"size:255;not null;" json:"token"`
	Active       bool   `gorm:"default:false" json:"active"`
	Verify       bool   `gorm:"default:false" json:"verify"`
	PermissionID uint64 `gorm:"not null" json:"permission_id"`
	LastLogin    string `gorm:"size:45;not null;" json:"last_login"`
    RemovedAt    string `gorm:"size:45;not null;" json:"removed_at"`
}

func (u *User) Validate() error {
	if IsEmpty(u.FirstName) || IsEmpty(u.LastName) || IsEmpty(u.Email) || IsEmpty(u.Password) || IsEmpty(u.Telephone) {
		return ErrEmptyFields	
	}
	
	if !IsEmail(u.Email) {
		return  ErrInvalidEmail
	}
	if !IsTelephone(u.Telephone) {
		return ErrInvalidTelephone
	}
	return nil
}

func (u *User) PutNewPasswordUserValidate() error {
	if IsEmpty(u.Email) || IsEmpty(u.Password) {
		return ErrEmptyFields	
	}
	
	if !IsEmail(u.Email) {
		return  ErrInvalidEmail
	}
	return nil
}

func (u *User) UserLoginValidate() error {
	if IsEmpty(u.Email) || IsEmpty(u.Password) {
		return ErrEmptyFields	
	}
	return nil
}

func IsEmpty(param string) bool {
	return param == ""
}

func IsEmail(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`).MatchString(value)
}

func IsTelephone(value string) bool {
	return  regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`).MatchString(value)
}
// Doesn't match password
func IsPassword(value string) bool {
	return regexp.MustCompile(`^[?=.*[0-9]][?=.*[!@#$%^&*]][a-zA-Z0-9!@#$%^&*]{8}$`).MatchString(value)
}
