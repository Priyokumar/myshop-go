package handler

import (
	"encoding/json"
	"log"
	"myshop/model"
	"myshop/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPurchasesHandler GetPurchasesHandler
func GetPurchasesHandler(ctx *gin.Context) {

	purchases := new([]model.Purchase)

	repo.FindPurchases(purchases)

	if purchases == nil || len(*purchases) <= 0 {
		ctx.JSON(404, "No purchase found.")
	} else {
		ctx.JSON(200, purchases)
	}

}

// GetPurchaseHandler GetPurchaseHandler
func GetPurchaseHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	purchase := new(model.Purchase)

	err := repo.FindPurchase(purchase, id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, purchase)

}

// AddPurchaseHandler AddPurchaseHandler
func AddPurchaseHandler(ctx *gin.Context) {

	body := ctx.Request.Body
	defer body.Close()

	purchase := new(model.Purchase)

	jsonDecodingError := json.NewDecoder(body).Decode(purchase)

	if jsonDecodingError != nil {
		log.Println(jsonDecodingError.Error())
		ctx.JSON(400, "Malformed request body")
		return
	}

	idStr := ctx.Param("id")

	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		purchase.ID = id
	}

	error := repo.AddPurchase(purchase)

	if error != nil {
		log.Println(error.Error())
		ctx.JSON(500, error.Error())
		return
	}

	ctx.JSON(200, "Ok")

}

// DeletePurchaseHandler DeletePurchaseHandler
func DeletePurchaseHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := repo.DeletePurchase(id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, "Deleted")
}
