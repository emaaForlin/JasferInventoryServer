package data

import (
	"time"

	"gorm.io/gorm"
)

func AddProduct(p *Product, db *gorm.DB) error {
	id, err := getNextID(db)
	if err != nil {
		p.ID = 1
	}
	p.ID = id
	p.CreatedAt = time.Now().UTC()

	prod := Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		SKU:         p.SKU,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   time.Date(1986, time.January, 1, 0, 0, 0, 0, time.Local),
	}
	if prod.ID != 0 && prod.Name != "" && prod.Price != 0 {
		err := db.Create(&prod).Error
		return err
	}
	return err
}
