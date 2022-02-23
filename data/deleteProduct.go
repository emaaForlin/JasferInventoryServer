package data

import "gorm.io/gorm"

func DeleteProduct(id int, db *gorm.DB) error {
	_, _, err := findProduct(id, db)
	if err != nil {
		return err
	}
	// this deletes permanently the entry
	db.Unscoped().Delete(&Product{}, id)
	return nil
}
