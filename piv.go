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

func drawImage(p image.Image) {
	bounds := p.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i := ansi.Index(p.At(x, y))
			fmt.Print(blocks[i])
		}
		fmt.Println()
	}
}

func scaledRectangle(r image.Rectangle, w int) image.Rectangle {
	size := r.Size()
	ratio := float64(size.X) / float64(size.Y)
	h := int(float64(w) / ratio)
	return image.Rect(0, 0, w, h)
}

func scaledImage(p image.Image, w int) image.Image {
	bounds := p.Bounds()
	dst := image.NewRGBA(scaledRectangle(bounds, w))
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), p, bounds, draw.Over, nil)
	return dst
}

func main() {
	width := flag.Int("width", 80, "output image width")
	flag.Parse()

	p, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	drawImage(scaledImage(p, *width))
}
