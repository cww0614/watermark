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
			cli.StringFlag{
				Name:  "color, c",
				Value: "blue",
				Usage: "Color for watermark text, name or #rrggbb",
			},
			cli.IntFlag{
				Name:  "vertical-spacing, e",
				Value: 40,
				Usage: "Vertical spacing between watermarks",
			},
			cli.StringFlag{
				Name:  "font, f",
				Value: "Courier",
				Usage: "Font for watermark text",
			},
			cli.IntFlag{
				Name:  "font-size, S",
				Value: 64,
				Usage: "Font size for watermark text",
			},
			cli.IntFlag{
				Name:  "horizontal-spacing, o",
				Value: 40,
				Usage: "Horizontal spacing between watermarks",
			},
			cli.Float64Flag{
				Name:  "scale, s",
				Value: 1.0,
				Usage: "Scale watermarks",
			},
			cli.Float64Flag{
				Name:  "transparency, t",
				Value: 0.05,
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

		color.A = uint8(c.Float64("transparency") * 255.0)

		waterMarker := &WaterMarker{
			Text:              text,
			HorizontalSpacing: c.Int("horizontal-spacing"),
			VerticalSpacing:   c.Int("vertical-spacing"),
			FontSize:          c.Int("font-size"),
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
