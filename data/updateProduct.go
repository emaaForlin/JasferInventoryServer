package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func UpdateProduct(id int, p *Product, db *gorm.DB) error {
	oldProd, pos, err := findProduct(id, db)
	if err != nil {
		return err
	}
	p.ID = id
	p.CreatedAt = oldProd.CreatedAt
	p.UpdatedAt = time.Now().UTC()

	if p.ID != 0 && p.Name != "" && p.Price != 0 && p.SKU != "" && p.CreatedAt != p.UpdatedAt {
		db.Save(&p)
		productList[pos] = p
		return nil
	}
	return fmt.Errorf("all values are needed")
}
