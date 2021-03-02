package qrcodeutil

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"myshop/model"
	"myshop/repo"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/math/fixed"
)

// QRCodeImagePath QRCodeImageBasePath
var QRCodeImagePath string = "images/qrcode"

// BarCodeImagePath BarCodeImagePath
var BarCodeImagePath string = "images/barcode"

// GenerateQRCode GenerateQRCode
func GenerateQRCode(content string, fileName string) (*int64, error) {

	filePath := QRCodeImagePath + "/item/" + fileName + ".png"

	err := qrcode.WriteFile(content, qrcode.Medium, 256, filePath)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	qrcode := new(model.QRCode)
	qrcode.Name = fileName
	qrcode.FilePath = filePath
	errInSaving := repo.SaveQRCode(qrcode)

	if errInSaving != nil {
		return nil, errInSaving
	}

	return &qrcode.ID, nil
}

// GenerateBarcode GenerateBarcode
func GenerateBarcode(content string, fileName string) (*int64, error) {

	fmt.Println("Generating code128 barcode for : ", content)

	bcodeInt, err := code128.Encode(content)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// scale to 250x20
	bcode, err := barcode.Scale(bcodeInt, 250, 70)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// Initialize the graphic context on an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 250, 50))

	// set background to white
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	addLabel(img, 50, 20, "prilax")

	// create a new blank image with white background
	newImg := imaging.New(250, 120, color.NRGBA{255, 255, 255, 255})

	//paste the codabar to new blank image
	newImg = imaging.Paste(newImg, bcode, image.Pt(0, 15))

	//paste the text to the new blank image
	newImg = imaging.Paste(newImg, img, image.Pt(50, 90))

	path := BarCodeImagePath + "/item"
	filePath := path + "/" + fileName + ".png"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeDir)
	}

	err = draw2dimg.SaveToPngFile(filePath, newImg)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	barcode := new(model.BarCode)
	barcode.Name = fileName
	barcode.FilePath = filePath
	errInSaving := repo.SaveBarCode(barcode)

	if errInSaving != nil {
		return nil, errInSaving
	}

	return &barcode.ID, nil
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{0, 0, 0, 255}

	var point fixed.Point26_6
	point.X = fixed.Int26_6(x * 64)
	point.Y = fixed.Int26_6(y * 64)

	newFont, _ := truetype.Parse(gobolditalic.TTF)

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(col),
		Face: truetype.NewFace(newFont, &truetype.Options{
			Size: 24,
			DPI:  72,
		}),
		Dot: point,
	}
	d.DrawString(label)
}
