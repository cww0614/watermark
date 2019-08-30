# watermarker

A command line tool to add watermark on images

## Demo

![demo](example.watermarked.png)

## Usage

```
Usage: watermarker [-hV] [-c=<color>] [-e=<verticalSpacing>] [-f=<fontName>]
                   [-o=<horizontalSpacing>] [-s=<scale>] [-S=<fontSize>]
                   [-t=<transparency>] TEXT FILE...
      TEXT                Watermark text
      FILE...             File(s) to process.
  -c, --color=<color>     Color for watermark text, name of #rrggbb (default:
                            blue)
  -e, --vertical-spacing=<verticalSpacing>
                          Vertical spacing between watermarks (default: 40)
  -f, --font=<fontName>   Font for watermark text (default: Arial)
  -h, --help              Show this help message and exit.
  -o, --horizontal-spacing=<horizontalSpacing>
                          Horizontal spacing between watermarks (default: 40)
  -s, --scale=<scale>     Scale watermarks (default: 1.0)
  -S, --font-size=<fontSize>
                          Font size for watermark text (default: 64)
  -t, --transparency=<transparency>
                          Transparency of watermark (default: 0.05)
  -V, --version           Print version information and exit.
```
