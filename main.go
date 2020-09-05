package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Config struct {
	Help            bool     `docopt:"-h,--help"`
	Version         bool     `docopt:"--version"`
	Width           int      `docopt:"-W,--width"`
	Height          int      `docopt:"-H,--height"`
	Column          int      `docopt:"-c,--column"`
	Row             int      `docopt:"-r,--row"`
	Pad             int      `docopt:"-p,--pad"`
	BackgroundColor string   `docopt:"-b,--background-color"`
	StrokeColor     string   `docopt:"-s,--stroke-color"`
	FillColor       string   `docopt:"-f,--fill-color"`
	LineWidth       float64  `docopt:"-l,--line-width"`
	Args            []string `docopt:"<args>"`
}

const version = `layoutimg v1.0.0
Copyright (c) 2020 jiro4989
Released under the MIT License.
https://github.com/jiro4989/layoutimg`

const usage = `layoutimg is TODO

Usage:
  layoutimg [options] <args>...
  layoutimg -h | --help
  layoutimg --version

Options:
  -h, --help                                   print this help
      --version                                print version
  -W, --width=<width>                          image rectangle width [default: 200]
  -H, --height=<height>                        image rectangle height [default: 200]
  -c, --column=<column>                        image tile columns count [default: 4]
  -r, --row=<row>                              image tile rows count [default: 4]
  -p, --pad=<pad>                              image tile padding width [default: 5]
  -b, --background-color=<background-color>    image background color [default: white]
  -s, --stroke-color=<stroke-color>            image stroke color [default: black]
  -f, --fill-color=<fill-color>                image file color [default: none]
  -l, --line-width=<line-width>                image line width [default: 2]

`

func main() {
	os.Exit(Main(os.Args[1:]))
}

func Main(args []string) int {
	opts, err := docopt.ParseArgs(usage, args, version)
	if err != nil {
		panic(err)
	}

	var config Config
	opts.Bind(&config)

	if config.Help {
		fmt.Println(usage)
		return 0
	}

	if config.Version {
		fmt.Println(version)
		return 0
	}

	dest := image.NewRGBA(image.Rect(0, 0, config.Width, config.Height))
	drawBackground(dest, colors[config.BackgroundColor])

	for _, xy := range config.Args {
		fs := strings.Split(xy, ",")
		x, err := strconv.Atoi(fs[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(fs[1])
		if err != nil {
			panic(err)
		}

		dp := drawParam{
			x:           x,
			y:           y,
			column:      config.Column,
			row:         config.Row,
			pad:         config.Pad,
			strokeColor: colors[config.StrokeColor],
			lineWidth:   config.LineWidth,
		}
		draw(dest, dp)
	}
	draw2dimg.SaveToPngFile("out.png", dest)

	return 0
}
