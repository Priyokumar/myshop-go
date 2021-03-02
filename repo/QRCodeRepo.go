package repo

import (
	"log"
	"myshop/database"
	"myshop/model"

	"gorm.io/gorm/clause"
)

// SaveQRCode SaveQRCode
func SaveQRCode(QRCode *model.QRCode) error {

	result := database.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(QRCode)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// SaveBarCode SaveBarCode
func SaveBarCode(barCode *model.BarCode) error {

	result := database.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(barCode)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// FindQRCode FindItem
func FindQRCode(qrCode *model.QRCode, id int64) error {

	result := database.DB.First(qrCode, "id = ?", id)

	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}
