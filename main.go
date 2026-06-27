package main

import (
	"log"

	"github.com/IosifSuzuki/go-watermark/watermark"
)

func main() {
	baseImg, err := watermark.LoadImage("examples/Listen.jpg")
	if err != nil {
		log.Fatalln("fail to open image file", err)
	}

	logoImg, err := watermark.LoadImage("examples/logo.png")
	if err != nil {
		log.Fatalln("fail to open logo file", err)
	}

	watermarkService := watermark.NewService()
	watermarkedImg := watermarkService.Apply(watermark.Request{
		Image:     baseImg,
		Logo:      logoImg,
		LogoAlpha: 0.5,
		LogoScale: 0.2,
		Placement: watermark.Bottom,
		Margins: watermark.Margins{
			Bottom: 45,
		},
	})

	if err := watermark.SaveImage("examples/Output.jpg", watermarkedImg, 100); err != nil {
		log.Fatalln("fail to save image", err)
	}
}
