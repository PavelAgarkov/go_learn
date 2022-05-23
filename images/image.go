package images

import (
	"fmt"
	"image"
)

const Im = 1
const im = 2

func Image() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
