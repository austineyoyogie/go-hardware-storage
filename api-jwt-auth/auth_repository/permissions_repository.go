package auth_repository

import (
	"time"

	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/models"
	"github.com/jinzhu/gorm"
)

type PermissionsRepository interface {
	Save(*models.Permission) (*models.Permission, error)
	Find(uint64) (*models.Permission, error)
	FindAll() ([]*models.Permission, error)
	Update(*models.Permission) error
	Delete(uint64) error
}

type permissionsRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionsRepository(db *gorm.DB) *permissionsRepositoryImpl {
	return &permissionsRepositoryImpl{db}
}

func (r *permissionsRepositoryImpl) Save(permission *models.Permission) (*models.Permission, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Permission{}).Create(permission).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return permission, tx.Commit().Error
}

func (r *permissionsRepositoryImpl) Find(permission_id uint64) (*models.Permission, error) {
	permission := &models.Permission{}
	err := r.db.Debug().Model(&models.Permission{}).Where("id = ?", permission_id).Find(permission).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Debug().Model(permission).Related(&permission.Users).Error
	return permission, err
}

func (r *permissionsRepositoryImpl) FindAll() ([]*models.Permission, error) {
	permissions := []*models.Permission{}
	err := r.db.Debug().Model(&models.Permission{}).Find(&permissions).Error
	return permissions, err
}

func (r *permissionsRepositoryImpl) Update(permission *models.Permission) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"role_name":  permission.RoleName,
		"updated_at": time.Now(),
	}

	err := tx.Debug().Model(&models.Permission{}).Where("id = ?", permission.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *permissionsRepositoryImpl) Delete(permission_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Permission{}).Where("id = ?", permission_id).Delete(&models.Permission{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
