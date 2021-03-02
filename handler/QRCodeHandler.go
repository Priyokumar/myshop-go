package handler

import (
	"fmt"
	"log"
	"myshop/model"
	"myshop/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QRCodeHandler QRCodeHandler
func QRCodeHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	qrCode := new(model.QRCode)

	err := repo.FindQRCode(qrCode, id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", qrCode.Name))
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.File(qrCode.FilePath)

}
