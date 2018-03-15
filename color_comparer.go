package hlc

import (
	"image/color"
)

type colorComparer interface {
	compare(c1, c2 color.Color) (percentage uint8) //Compute the same rate of 2 Colors
}

// type Test color.Color

// func (src Test) Compare(dest color.Color) (percentage uint8) {
// 	return 10
// }

// type defaultColorComparer struct {
// }

// func (cp defaultColorComparer) compare(c1, c2 color.Color) (percentage uint8) {
// 	r1, g1, b1, _ := convertColorTo8BitRGBA(c1)
// 	r2, g2, b2, _ := convertColorTo8BitRGBA(c2)

// 	var sum = int(math.Abs(float64(r1)-float64(r2)) + math.Abs(float64(g1)-float64(g2)) + math.Abs(float64(b1)-float64(b2)))
// }
