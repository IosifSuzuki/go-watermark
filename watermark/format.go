package watermark

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	_ "golang.org/x/image/webp"
)

type Format string

const (
	JPEG Format = "jpeg"
	PNG  Format = "png"
	WEBP Format = "webp"
)

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	img, _, err := image.Decode(file)
	return img, err
}

func SaveImage(path string, img image.Image, quality int) error {
	format, err := formatFromPath(path)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	return Encode(file, img, format, quality)
}

func Encode(w io.Writer, img image.Image, format Format, quality int) error {
	switch format {
	case JPEG:
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	case PNG:
		return png.Encode(w, img)
	case WEBP:
		return fmt.Errorf("webp encoding is not supported")
	default:
		return fmt.Errorf("unsupported image format %q", format)
	}
}

func formatFromPath(path string) (Format, error) {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".jpg", ".jpeg":
		return JPEG, nil
	case ".png":
		return PNG, nil
	case ".webp":
		return WEBP, nil
	default:
		return "", fmt.Errorf("unsupported image extension %q", filepath.Ext(path))
	}
}
