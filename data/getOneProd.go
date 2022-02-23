package data

import "gorm.io/gorm"

func GetOneProduct(db *gorm.DB, id int) ([]*Product, error) {
	db.First(&productList, id)
	return productList, db.Error
}
