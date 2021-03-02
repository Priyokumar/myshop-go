package repo

import (
	"log"
	"myshop/database"
	"myshop/model"

	"gorm.io/gorm/clause"
)

// AddProduct AddProduct
func AddProduct(product *model.Product) error {

	result := database.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(product)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindProducts FindProducts
func FindProducts(products *[]model.Product) error {

	result := database.DB.Find(&products)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindProductsCount FindProductsCount
func FindProductsCount() (int64, error) {

	var count int64
	result := database.DB.Model(&model.Product{}).Count(&count)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return count, result.Error
	}

	return count, nil
}

// FindProduct FindProduct
func FindProduct(product *model.Product, id int64) error {

	result := database.DB.First(product, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindProductByBarcode FindProductByBarcode
func FindProductByBarcode(product *model.Product, barcode string) error {

	result := database.DB.First(product, "barcode = ?", barcode)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// DeleteProduct DeleteProduct
func DeleteProduct(id int64) error {

	result := database.DB.Delete(&model.Product{}, id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}
