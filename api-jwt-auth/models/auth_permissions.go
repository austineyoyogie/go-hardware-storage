package models

import "errors"

type Permission struct {
	Model
	RoleName string `gorm:"size:100;not null;" json:"role_name" validate:"required"`
	Users []*User `gorm:"foreignkey:PermissionID" json:"users"`
	RemovedAt string `gorm:"size:100;not null;" json:"removed_at"`
}

var (
	ErrPermissionEmptyRoleName = errors.New("permission.role name can't be empty")
)

func (p *Permission) Validate() error {
	if p.RoleName == "" {
		return ErrPermissionEmptyRoleName
	}
	return nil
}


// github.com/go-playground/validator
// https://www.golangprograms.com/go-struct-and-field-validation-examples.html

// https://github.com/qor/validations/blob/master/validation_test.go

// https://itnext.io/validating-struct-map-form-in-go-language-1f819b8596c7