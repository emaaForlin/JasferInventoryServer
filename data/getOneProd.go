package data

import "gorm.io/gorm"

func GetOneProduct(db *gorm.DB, id int) ([]*Product, error) {
	err := db.First(&productList, id).Error
	return productList, err
}
