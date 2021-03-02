package repo

import (
	"log"
	"myshop/database"
	"myshop/model"

	"gorm.io/gorm/clause"
)

// AddItem AddItem
func AddItem(item *model.Item) error {

	result := database.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(item)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindItems FindItems
func FindItems(items *[]model.Item) (*[]model.Item, error) {

	result := database.DB.Find(&items)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return items, nil
}

// FindItemsCount FindItemsCount
func FindItemsCount() (int64, error) {

	var count int64
	result := database.DB.Model(&model.Item{}).Count(&count)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return count, result.Error
	}

	return count, nil
}

// FindItem FindItem
func FindItem(item *model.Item, id int64) error {

	result := database.DB.First(item, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// DeleteItem DeleteItem
func DeleteItem(id int64) error {

	result := database.DB.Delete(&model.Item{}, id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}
