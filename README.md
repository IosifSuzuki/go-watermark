# watermark-go

Small Go package for applying an image logo watermark to another image.

The watermark package accepts decoded `image.Image` values, scales the logo, applies global alpha through a mask, and places it at a requested position.

## Install

For local development in this module:

```sh
go get watermark-go/watermark
```

When published as a remote module, replace `watermark-go` with the repository module path, for example:

```go
import "github.com/yourname/watermark-go/watermark"
```

## Basic Usage

```go
package main

import (
	"log"

	"watermark-go/watermark"
)

func main() {
	baseImg, err := watermark.LoadImage("examples/Listen.jpg")
	if err != nil {
		log.Fatal(err)
	}

	logoImg, err := watermark.LoadImage("examples/logo.png")
	if err != nil {
		log.Fatal(err)
	}

	service := watermark.NewService()
	result := service.Apply(watermark.Request{
		Image:     baseImg,
		Logo:      logoImg,
		LogoAlpha: 0.4,
		LogoScale: 0.2,
		Placement: watermark.BottomRight,
		Margins: watermark.Margins{
			Right:  45,
			Bottom: 45,
		},
	})

	if err := watermark.SaveImage("examples/output.jpg", result, 90); err != nil {
		log.Fatal(err)
	}
}
```

## Request Options

```go
type Request struct {
	Image     image.Image
	Logo      image.Image
	LogoAlpha float64
	LogoScale float64
	Placement Placement
	Margins   Margins
}
```

`LogoAlpha` controls logo opacity:

```go
LogoAlpha: 0.25 // 25%
LogoAlpha: 0.50 // 50%
LogoAlpha: 1.00 // 100%
```

`LogoScale` controls the logo size relative to the smaller side of the base image:

```go
LogoScale: 0.1 // smaller
LogoScale: 0.2 // default example size
LogoScale: 0.4 // larger
```

`Margins` controls distance from the selected edge:

```go
Margins: watermark.Margins{
	Top:    20,
	Right:  20,
	Bottom: 20,
	Left:   20,
}
```

## Placements

```go
watermark.Top
watermark.Right
watermark.Bottom
watermark.Left
watermark.TopLeft
watermark.TopRight
watermark.BottomLeft
watermark.BottomRight
watermark.Center
```

Single-edge placements are centered on the opposite axis. For example, `watermark.Top` is horizontally centered and uses `Margins.Top`.

## Supported Formats

Loading:

```text
.jpg
.jpeg
.png
.webp
```

Saving:

```text
.jpg
.jpeg
.png
```

WebP decoding is supported through `golang.org/x/image/webp`. WebP encoding is not supported by the current implementation.

## Lower-Level Encoding

Use `Encode` when writing to your own `io.Writer`:

```go
err := watermark.Encode(writer, img, watermark.JPEG, 90)
```

For PNG, the quality value is ignored:

```go
err := watermark.Encode(writer, img, watermark.PNG, 0)
```
