package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func outputFilename(input string) string {
	lastDot := strings.LastIndex(input, ".")
	return input[:lastDot] + ".watermarked" + input[lastDot:]
}

func main() {
	app := &cli.App{
		Name:      "watermark",
		Usage:     "add watermark on images",
		UsageText: "watermark [OPTIONS] TEXT FILE ...",
		HideHelp:  true,
		Version:   "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "color, c",
				Value: "blue",
				Usage: "Color for watermark text, name or #rrggbb",
			},
			&cli.Float64Flag{
				Name:  "vertical-spacing, e",
				Value: 40.0,
				Usage: "Vertical spacing between watermarks",
			},
			&cli.IntFlag{
				Name:  "output-dpi, d",
				Value: 72,
				Usage: "DPI of output image",
			},
			&cli.StringFlag{
				Name:  "font, f",
				Value: "Courier",
				Usage: "Font for watermark text",
			},
			&cli.Float64Flag{
				Name:  "font-size, S",
				Value: 64.0,
				Usage: "Font size for watermark text",
			},
			&cli.Float64Flag{
				Name:  "horizontal-spacing, o",
				Value: 40.0,
				Usage: "Horizontal spacing between watermarks",
			},
			&cli.Float64Flag{
				Name:  "scale, s",
				Value: 1.0,
				Usage: "Scale watermarks",
			},
			&cli.Float64Flag{
				Name:  "transparency, t",
				Value: 0.90,
				Usage: "Transparency of watermark",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() < 2 {
			return fmt.Errorf("not enough arguments")
		}

		text := c.Args().Get(0)

		color, err := parseColor(c.String("color"))
		if err != nil {
			return err
		}

		color.A = uint8((1 - c.Float64("transparency")) * 255.0)

		scale := c.Float64("scale")

		waterMarker := &WaterMarker{
			Text:              text,
			HorizontalSpacing: c.Float64("horizontal-spacing") * scale,
			VerticalSpacing:   c.Float64("vertical-spacing") * scale,
			FontSize:          c.Float64("font-size") * scale,
			OutputDPI:         c.Int("output-dpi"),
			FontName:          c.String("font"),
			Color:             color,
		}

		for i := 1; i < c.NArg(); i++ {
			file := c.Args().Get(i)

			output := outputFilename(file)
			err := waterMarker.Mark(file, output)
			if err != nil {
				return err
			}

			fmt.Println(file, "->", output)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("error:", err)
	}
}
