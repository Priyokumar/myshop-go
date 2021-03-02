package config

import (
	"myshop/handler"

	"github.com/gin-gonic/gin"
)

// SetupRoutes SetupRoutes
func SetupRoutes(routeEngine *gin.Engine) {

	v1 := routeEngine.Group("/v1/api")

	// For custom products
	v1.GET("/custom/products/count", handler.GetItemsCountHandler)
	v1.POST("/custom/products/bulk", handler.AddBulkItemsHandler)
	v1.PUT("/custom/products/bulk", handler.UpdateItemsHandler)
	v1.PUT("/custom/products", handler.UpdateItemsHandler)
	v1.GET("/custom/products", handler.GetItemsHandler)
	v1.POST("/custom/products", handler.AddItemHandler)
	v1.GET("/custom/product/:id", handler.GetItemHandler)
	v1.PUT("/custom/product/:id", handler.AddItemHandler)

	// For normal products
	v1.GET("/products", handler.GetProductsHandler)
	v1.POST("/products", handler.AddProductHandler)
	v1.GET("/products/count/", handler.GetProductByBarcodeHandler)
	v1.GET("/products/:id", handler.GetProductHandler)
	v1.PUT("/product/:id", handler.AddProductHandler)
	v1.DELETE("/product/:id", handler.DeleteProductHandler)
	v1.GET("/product/barcode/:barcode", handler.GetProductByBarcodeHandler)

	// Qrcode reader
	v1.GET("/images/qrcode/:id", handler.QRCodeHandler)

	// For Purchase

	v1.GET("/purchases", handler.GetPurchasesHandler)
	v1.POST("/purchases", handler.AddPurchaseHandler)
	v1.GET("/purchases/:id", handler.GetPurchaseHandler)

}
