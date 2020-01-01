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

	_ "golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var palette = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // Black
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // Red
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // Green
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // Yellow
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // Blue
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // Magenta
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // Cyan
	color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
	color.RGBA{0x00, 0x00, 0x00, 0x00}, // Transparent
}

const end = "\033[0m"

var blocks = []string{
	"\033[40m ",
	"\033[41m ",
	"\033[42m ",
	"\033[43m ",
	"\033[44m ",
	"\033[45m ",
	"\033[46m ",
	"\033[47m ",
	end + " ",
}

func drawImage(p image.Image) {
	bounds := p.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i := palette.Index(p.At(x, y))
			fmt.Print(blocks[i])
		}
		fmt.Println(end)
	}
}

func scaledRectangle(r image.Rectangle, w int, ratio float64) image.Rectangle {
	size := r.Size()
	if w <= 0 {
		w = size.X
	}
	h := (ratio * float64(w*size.Y)) / float64(size.X)
	return image.Rect(0, 0, w, int(h))
}

func scaledImage(p image.Image, r image.Rectangle) image.Image {
	dst := image.NewRGBA(r)
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), p, p.Bounds(), draw.Over, nil)
	return dst
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage of piv:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Image data is read from standard input.")
	}
	w := flag.Int("w", 80, "output image width (native width if 0)")
	ratio := flag.Float64("r", 0.5, "character width-to-height ratio")
	flag.Parse()

	p, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	r := scaledRectangle(p.Bounds(), *w, *ratio)
	drawImage(scaledImage(p, r))
}
