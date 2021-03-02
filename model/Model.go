package model

import (
	"time"

	"gorm.io/gorm"
)

// Item Item
type Item struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Unit          string    `json:"unit"`
	Price         float64   `json:"price"`
	Quantity      int64     `json:"quantity"`
	QRCodeURL     string    `json:"qrCodeURL"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" gorm:"autoCreateTime"`
}

// Product Product
type Product struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Barcode       string    `json:"barcode"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	Stock         int64     `json:"stock"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" gorm:"autoCreateTime"`
}

// Purchase Purchase
type Purchase struct {
	gorm.Model
	ID            int64           `json:"id" gorm:"primaryKey"`
	CustomerName  string          `json:"customerName"`
	MobileNo      string          `json:"mobileNo"`
	TotalPrice    float64         `json:"totalPrice"`
	TotalItem     int64           `json:"totalItem"`
	Items         *[]PurchaseItem `json:"items" gorm:"foreignkey:PurchaseID;references:ID"`
	CreatedAt     time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	LastUpdatedAt time.Time       `json:"lastUpdatedAt" gorm:"autoCreateTime"`
}

// PurchaseItem PurchaseItem
type PurchaseItem struct {
	gorm.Model
	ID            int64     `json:"id"`
	Barcode       string    `json:"barcode"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	Quantity      int64     `json:"quantity"`
	PurchaseID    string    `json:"purchaseId"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" gorm:"autoCreateTime"`
}

// QRCode QRCode
type QRCode struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	FilePath      string    `json:"filePath"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" gorm:"autoCreateTime"`
}

// BarCode QRCode
type BarCode struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	FilePath      string    `json:"filePath"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt" gorm:"autoCreateTime"`
}
