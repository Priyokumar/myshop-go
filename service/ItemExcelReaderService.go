package service

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"myshop/model"
	"myshop/repo"
	"myshop/util/qrcodeutil"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ReadExcelNSave ReadExcelNSave
func ReadExcelNSave(file multipart.File) error {

	raw, readExcelErr := readExcel(file)

	if readExcelErr != nil {
		return readExcelErr
	}

	saveErr := saveAllItems(raw)

	if saveErr != nil {
		return saveErr
	}

	return nil
}

// ReadExcel ReadExcel
func readExcel(file multipart.File) ([][]string, error) {

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	workBook, err := excelize.OpenReader(buf)

	if err != nil {
		println(err.Error())
		return nil, err
	}

	allRows := make([][]string, 0)
	for _, sheetName := range workBook.GetSheetMap() {
		rows := workBook.GetRows(sheetName)
		allRows = append(allRows, rows...)
		log.Println(rows)
	}

	return allRows, nil
}

func saveAllItems(raw [][]string) error {

	length := len(raw)

	for i := 1; i < length; i++ {
		saveItem(raw[i])
	}
	return nil
}

func saveItem(itemRaw []string) error {

	item := new(model.Item)

	item.Name = itemRaw[0]
	item.Category = itemRaw[1]
	item.Unit = itemRaw[2]
	price, priceErr := strconv.ParseFloat(itemRaw[3], 64)
	if priceErr != nil {
		return priceErr
	}
	item.Price = price
	quantity, quantityErr := strconv.ParseInt(itemRaw[4], 10, 64)
	if quantityErr != nil {
		return priceErr
	}
	item.Quantity = quantity

	saveErr := repo.AddItem(item)

	if saveErr != nil {
		return saveErr
	}

	idStr := strconv.FormatInt(item.ID, 10)

	qrCodeID, errInQRCode := qrcodeutil.GenerateBarcode(idStr, strconv.FormatInt(item.ID, 10))

	if errInQRCode != nil {
		return errInQRCode
	}

	qrCodeIDStr := strconv.FormatInt(*qrCodeID, 10)

	qrCodeURL := "/v1/api/images/qrcode/" + qrCodeIDStr

	item.QRCodeURL = qrCodeURL

	updatedErr := repo.AddItem(item)

	if updatedErr != nil {
		return saveErr
	}

	return nil
}
