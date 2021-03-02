package handler

import (
	"encoding/json"
	"log"
	"myshop/model"
	"myshop/repo"
	"myshop/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetItemsHandler GetItemsHandler
func GetItemsHandler(ctx *gin.Context) {

	items := make([]model.Item, 0)

	repo.FindItems(&items)

	ctx.JSON(200, items)

}

// GetItemsCountHandler GetItemCountHandler
func GetItemsCountHandler(ctx *gin.Context) {

	count, err := repo.FindItemsCount()

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, count)

}

// GetItemHandler GetItemHandler
func GetItemHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	item := new(model.Item)

	err := repo.FindItem(item, id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, item)

}

// AddItemHandler AddItemHandler
func AddItemHandler(ctx *gin.Context) {

	body := ctx.Request.Body
	defer body.Close()

	item := new(model.Item)

	jsonDecodingError := json.NewDecoder(body).Decode(item)

	if jsonDecodingError != nil {
		log.Println(jsonDecodingError.Error())
		ctx.JSON(400, "Malformed request body")
		return
	}

	idStr := ctx.Param("id")

	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		item.ID = id
	}

	error := repo.AddItem(item)

	if error != nil {
		log.Println(error.Error())
		ctx.JSON(500, error.Error())
		return
	}

	ctx.JSON(200, "Ok")

}

// UpdateItemsHandler UpdateItemHandler
func UpdateItemsHandler(ctx *gin.Context) {

	body := ctx.Request.Body
	defer body.Close()

	items := new([]model.Item)

	jsonDecodingError := json.NewDecoder(body).Decode(items)

	if jsonDecodingError != nil {
		log.Println(jsonDecodingError.Error())
		ctx.JSON(400, "Malformed request body")
		return
	}

	for _, item := range *items {
		error := repo.AddItem(&item)
		if error != nil {
			log.Println(error.Error())
			ctx.JSON(500, error.Error())
			return
		}
	}

	ctx.JSON(200, "Ok")

}

// AddBulkItemsHandler AddBulkItemsHandler
func AddBulkItemsHandler(ctx *gin.Context) {

	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	if len(files) <= 0 {
		log.Println(files)
		ctx.JSON(400, "No file found.")
		return
	}

	file, fileOpenError := files[0].Open()

	if fileOpenError != nil {
		ctx.JSON(500, fileOpenError.Error())
		return
	}

	readWriteErr := service.ReadExcelNSave(file)

	if readWriteErr != nil {
		ctx.JSON(500, readWriteErr.Error())
		return
	}

	ctx.JSON(200, "ok")

}

// DeleteItemHandler DeleteItemHandler
func DeleteItemHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := repo.DeleteItem(id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, "Deleted")

}
