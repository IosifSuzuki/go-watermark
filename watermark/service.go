package watermark

import (
	"image"

	"golang.org/x/image/draw"
)

type Request struct {
	Image     image.Image
	Logo      image.Image
	LogoAlpha float64
	LogoScale float64
	Placement Placement
	Margins   Margins
}

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Apply(request Request) image.Image {
	logoAspectRatio := float64(request.Logo.Bounds().Dx()) / float64(request.Logo.Bounds().Dy())
	bounds := request.Image.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, request.Image, image.Point{}, draw.Src)

	logoScale := max(0.1, request.LogoScale)

	maxLogoDimension := int(min(float64(rgba.Bounds().Dx())*logoScale, float64(rgba.Bounds().Dy())*logoScale))
	logoSize := size{
		width:  maxLogoDimension,
		height: maxLogoDimension,
	}
	if logoAspectRatio > 1 {
		logoSize.height = int(float64(logoSize.height) / logoAspectRatio)
	} else {
		logoSize.width = int(float64(logoSize.width) * logoAspectRatio)
	}

	logoRectangle := s.logoRectangle(rgba.Rect, logoSize, request.Margins, request.Placement)

	scaledLogo := image.NewRGBA(image.Rect(0, 0, logoSize.width, logoSize.height))
	draw.NearestNeighbor.Scale(scaledLogo, scaledLogo.Bounds(), request.Logo, request.Logo.Bounds(), draw.Src, nil)

	mask := newAlphaMask(request.LogoAlpha)
	draw.DrawMask(rgba, logoRectangle, scaledLogo, image.Point{}, mask, image.Point{}, draw.Over)

	return rgba
}

func (s Service) logoRectangle(bounds image.Rectangle, logoSize size, margins Margins, placement Placement) image.Rectangle {
	var min image.Point

	switch placement {
	case Top:
		min = image.Point{
			X: bounds.Min.X + (bounds.Dx()-logoSize.width)/2,
			Y: bounds.Min.Y + margins.Top,
		}
	case Right:
		min = image.Point{
			X: bounds.Dx() - logoSize.width - margins.Right,
			Y: bounds.Min.Y + (bounds.Dy()-logoSize.height)/2,
		}
	case Bottom:
		min = image.Point{
			X: bounds.Min.X + (bounds.Dx()-logoSize.width)/2,
			Y: bounds.Dy() - logoSize.height - margins.Bottom,
		}
	case Left:
		min = image.Point{
			X: bounds.Min.X + margins.Left,
			Y: bounds.Min.Y + (bounds.Dy()-logoSize.height)/2,
		}
	case TopLeft:
		min = image.Point{
			X: bounds.Min.X + margins.Left,
			Y: bounds.Min.Y + margins.Top,
		}
	case TopRight:
		min = image.Point{
			X: bounds.Dx() - logoSize.width - margins.Right,
			Y: bounds.Min.Y + margins.Top,
		}
	case BottomLeft:
		min = image.Point{
			X: bounds.Min.X + margins.Left,
			Y: bounds.Dy() - logoSize.height - margins.Bottom,
		}
	case Center:
		min = image.Point{
			X: bounds.Min.X + (bounds.Dx()-logoSize.width)/2,
			Y: bounds.Min.Y + (bounds.Dy()-logoSize.height)/2,
		}
	default:
		min = image.Point{
			X: bounds.Dx() - logoSize.width - margins.Right,
			Y: bounds.Dy() - logoSize.height - margins.Bottom,
		}
	}

	return image.Rectangle{
		Min: min,
		Max: image.Point{
			X: min.X + logoSize.width,
			Y: min.Y + logoSize.height,
		},
	}
}
