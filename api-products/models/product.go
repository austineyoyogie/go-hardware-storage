package models

import "errors"

type ProductStatus uint8

const (
	ProductStatus_Unavailable = 0
	ProductStatus_Available   = 1
)

type Product struct {
	Model
	SUK 	 string 	   `gorm:"size:50;not null;unique" json:"suk"`
	Name 	 string 	   `gorm:"size:512;not null;unique" json:"name"`
	Price 	 float64 	   `gorm:"type:decimal(10,2);not null;default:0.0" json:"price"`
	Quantity uint16 	   `gorm:"default:0;unsigned" json:"quantity"`
	Image 	 string 	   `gorm:"size:512;not null;" json:"image"`
	Status   ProductStatus `gorm:"char(1);default:0" json:"status"`
	CategoryID uint64	   `gorm:"not null" json:"category_id"`
}

var (
	ErrProductEmptyName = errors.New("product.name can't be empty")
)

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductEmptyName
	}
	return nil
}

/*
{
    "name": "GTX 1660ti",
    "price": 1599.88,
    "quantity": 50,
    "status": 1,
    "category_id": 1
}
*/

// select products.name, categories.description from products inner join categories on categories.id = products.category_id;