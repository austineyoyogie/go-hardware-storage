package models

type Group struct {
	Model
	Name         string `gorm:"size:100;not null;" json:"name"`
	PermissionID uint64 `gorm:"not null" json:"permission_id"`
}