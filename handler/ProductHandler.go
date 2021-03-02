package handler

import (
	"encoding/json"
	"log"
	"myshop/model"
	"myshop/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProductsHandler GetProductsHandler
func GetProductsHandler(ctx *gin.Context) {

	products := new([]model.Product)

	repo.FindProducts(products)

	ctx.JSON(200, products)

}

// GetProductsCountHandler GetProductsCountHandler
func GetProductsCountHandler(ctx *gin.Context) {

	count, err := repo.FindProductsCount()

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, count)

}

// GetProductHandler GetProductHandler
func GetProductHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	product := new(model.Product)

	err := repo.FindProduct(product, id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, product)

}

// GetProductByBarcodeHandler GetProductByBarcodeHandler
func GetProductByBarcodeHandler(ctx *gin.Context) {

	barcode := ctx.Param("barcode")

	product := new(model.Product)

	err := repo.FindProductByBarcode(product, barcode)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, product)
}

// AddProductHandler AddProductHandler
func AddProductHandler(ctx *gin.Context) {

	body := ctx.Request.Body
	defer body.Close()

	product := new(model.Product)

	jsonDecodingError := json.NewDecoder(body).Decode(product)

	if jsonDecodingError != nil {
		log.Println(jsonDecodingError.Error())
		ctx.JSON(400, "Malformed request body")
		return
	}

	idStr := ctx.Param("id")

	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		product.ID = id
	}

	error := repo.AddProduct(product)

	if error != nil {
		log.Println(error.Error())
		ctx.JSON(500, error.Error())
		return
	}

	ctx.JSON(200, "Ok")

}

// UpdateProductsHandler UpdateProductsHandler
func UpdateProductsHandler(ctx *gin.Context) {

	body := ctx.Request.Body
	defer body.Close()

	products := new([]model.Product)

	jsonDecodingError := json.NewDecoder(body).Decode(products)

	if jsonDecodingError != nil {
		log.Println(jsonDecodingError.Error())
		ctx.JSON(400, "Malformed request body")
		return
	}

	for _, product := range *products {
		error := repo.AddProduct(&product)
		if error != nil {
			log.Println(error.Error())
			ctx.JSON(500, error.Error())
			return
		}
	}

	ctx.JSON(200, "Ok")

}

// DeleteProductHandler DeleteProductHandler
func DeleteProductHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := repo.DeleteProduct(id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, "Deleted")
}
