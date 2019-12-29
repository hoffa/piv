package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

var ansi = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

var blocks = []string{
	"\033[40m \033[0m",
	"\033[41m \033[0m",
	"\033[42m \033[0m",
	"\033[43m \033[0m",
	"\033[44m \033[0m",
	"\033[45m \033[0m",
	"\033[46m \033[0m",
	"\033[47m \033[0m",
}

func main() {
	width := flag.Int("width", 80, "width")
	flag.Parse()
	m, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	b := m.Bounds()
	ratio := float64(b.Max.X) / float64(b.Max.Y)
	dst := image.NewRGBA(image.Rect(0, 0, *width, int(float64(*width)/ratio)))
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), m, m.Bounds(), draw.Over, nil)
	bounds := dst.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			fmt.Print(blocks[ansi.Index(dst.At(x, y))])
		}
		fmt.Println()
	}
}
