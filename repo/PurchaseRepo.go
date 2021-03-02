package repo

import (
	"log"
	"myshop/database"
	"myshop/model"

	"gorm.io/gorm/clause"
)

// AddPurchase AddPurchase
func AddPurchase(purchase *model.Purchase) error {

	result := database.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(purchase)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindPurchases FindPurchases
func FindPurchases(purchases *[]model.Purchase) error {

	result := database.DB.Preload("Items").Find(&purchases)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindPurchasesCount FindPurchasesCount
func FindPurchasesCount() (int64, error) {

	var count int64
	result := database.DB.Model(&model.Purchase{}).Count(&count)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return count, result.Error
	}

	return count, nil
}

// FindPurchase FindPurchase
func FindPurchase(purchase *model.Purchase, id int64) error {

	result := database.DB.First(purchase, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// DeletePurchase DeletePurchase
func DeletePurchase(id int64) error {

	result := database.DB.Delete(&model.Purchase{}, id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}
