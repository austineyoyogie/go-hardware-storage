package product_repository

import (
	"time"

	"github.com/austineyoyogie/go-hardware-store/api-products/models"
	"github.com/jinzhu/gorm"
)
// Create as Service interface
type CategoriesRepository interface {
	Save(*models.Category) (*models.Category, error)
	Find(uint64) (*models.Category, error)
	FindAll() ([]*models.Category, error)
	Update(*models.Category) error
	Delete(uint64) error
}

type categoriesRepositoryImpl struct {
	db *gorm.DB
}
// Create as Service Impl
// https://www.youtube.com/watch?v=lrIR3caBlaM
// https://www.youtube.com/watch?v=vDIAwtGU9LE
// https://www.youtube.com/watch?v=dHSs4XSkCdk&list=PL5dTjWUk_cPY7Q2VTnMbbl8n-H4YDI5wF&index=6

func NewCategoriesRepository(db *gorm.DB) *categoriesRepositoryImpl {
	return &categoriesRepositoryImpl{db}
}

func (r *categoriesRepositoryImpl) Save(category *models.Category) (*models.Category, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Category{}).Create(category).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return category, tx.Commit().Error
}

func (r *categoriesRepositoryImpl) Find(category_id uint64) (*models.Category, error) {
	category := &models.Category{}
	err := r.db.Debug().Model(&models.Category{}).Where("id = ?", category_id).Find(category).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Debug().Model(category).Related(&category.Products).Error
	return category, err
}

func (r *categoriesRepositoryImpl) FindAll() ([]*models.Category, error) {
	categories := []*models.Category{}
	err := r.db.Debug().Model(&models.Category{}).Find(&categories).Error
	return categories, err
}

func (r *categoriesRepositoryImpl) Update(category *models.Category) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"description": category.Description,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.Category{}).Where("id = ?", category.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *categoriesRepositoryImpl) Delete(category_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Category{}).Where("id = ?", category_id).Delete(&models.Category{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

