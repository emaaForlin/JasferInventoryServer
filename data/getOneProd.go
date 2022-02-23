package data

import "gorm.io/gorm"

func GetOneProduct(db *gorm.DB, id int) []*Product {
	db.First(&productList, id)
	return productList
}
