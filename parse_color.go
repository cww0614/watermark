package main

import (
	"fmt"
	"image/color"
	"strings"
)

var colorMap = map[string]color.RGBA{
	"red":  color.RGBA{255, 0, 0, 0},
	"gree": color.RGBA{0, 255, 0, 0},
	"blue": color.RGBA{0, 0, 255, 0},
}

func parseColor(name string) (color.RGBA, error) {
	name = strings.ToLower(name)

	if len(name) == 7 && name[0] == '#' {
		return parseHex(name)
	} else {
		return parseName(name)
	}
}

func hexStrToInt(str string) (uint8, error) {
	value := uint8(0)
	for i := len(str) - 1; i >= 0; i-- {
		c := str[i]
		v := byte(0)

		switch {
		case '0' <= c && c <= '9':
			v = c - '0'
		case 'a' <= c && c <= 'f':
			v = c - 'a' + 10
		default:
			return 0, fmt.Errorf("invalid char in color str")
		}

		value = value * 16
		value += uint8(v)
	}

	return value, nil
}

func parseHex(hex string) (color.RGBA, error) {
	r, err := hexStrToInt(hex[1:3])
	if err != nil {
		return color.RGBA{}, err
	}

	g, err := hexStrToInt(hex[3:5])
	if err != nil {
		return color.RGBA{}, err
	}

	b, err := hexStrToInt(hex[5:7])
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{r, g, b, 0}, nil
}

func parseName(name string) (color.RGBA, error) {
	c, ok := colorMap[name]
	if ok {
		return c, nil
	} else {
		return color.RGBA{}, fmt.Errorf("cannot find color with name")
	}
}
