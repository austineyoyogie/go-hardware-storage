package models

import "errors"

type Category struct {
	Model
	SUK 	    string 	`gorm:"size:50;not null;unique" json:"suk"`
	Description string  `gorm:"size:512;not null; unique" json:"description"`
	Products []*Product `gorm:"foreignkey:CategoryID" json:"products"`
}

var (
	ErrCategoryEmptyDescription = errors.New("category.description can't be empty")
)

func (c *Category) Validate() error {
	if c.Description == "" {
		return ErrCategoryEmptyDescription
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