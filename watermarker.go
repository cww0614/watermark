package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"path/filepath"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

type WaterMarker struct {
	Text              string
	HorizontalSpacing int
	VerticalSpacing   int
	FontSize          int
	FontName          string
	Color             color.Color
}

func l(v int) vg.Length {
	return vg.Length(float64(v))
}

func (w *WaterMarker) Mark(inputFilename, outputFilename string) error {
	input, err := os.Open(inputFilename)
	if err != nil {
		return err
	}

	defer input.Close()

	output, err := os.Create(outputFilename)
	if err != nil {
		return err
	}

	defer output.Close()

	ext := filepath.Ext(outputFilename)

	img, _, err := image.Decode(input)
	if err != nil {
		return err
	}

	return w.mark(img, ext, output)
}

func (w *WaterMarker) mark(img image.Image, format string, out io.Writer) error {
	bounds := img.Bounds()

	width := bounds.Max.X
	height := bounds.Max.Y

	c := vgimg.New(l(width), l(height))
	c.DrawImage(vg.Rectangle{Max: vg.Point{X: l(width), Y: l(height)}}, img)

	c.SetColor(w.Color)

	fontStyle, err := vg.MakeFont(w.FontName, l(w.FontSize))
	if err != nil {
		return err
	}

	textWidth := int(fontStyle.Width(w.Text))
	textHeight := w.FontSize

	textBoxWidth := textWidth + w.HorizontalSpacing*2
	textBoxHeight := textHeight + w.VerticalSpacing*2

	c.Rotate(-math.Pi / 4)

	xOffsetMin := int(-float64(height) / math.Sqrt(2))
	yOffsetMax := int(float64(width) * math.Sqrt(2))

	for xOffset := xOffsetMin; xOffset < width; xOffset += textBoxWidth {
		for yOffset := 0; yOffset < yOffsetMax; yOffset += textBoxHeight {
			c.FillString(fontStyle, vg.Point{
				X: l(xOffset),
				Y: l(yOffset),
			}, w.Text)
		}
	}

	return writeCanvas(c, format, out)
}

func writeCanvas(c *vgimg.Canvas, format string, out io.Writer) error {
	switch format {
	case ".jpeg", ".jpg":
		_, err := vgimg.JpegCanvas{Canvas: c}.WriteTo(out)
		return err
	case ".png":
		_, err := vgimg.PngCanvas{Canvas: c}.WriteTo(out)
		return err
	default:
		return fmt.Errorf("unsupported file format")
	}
}
