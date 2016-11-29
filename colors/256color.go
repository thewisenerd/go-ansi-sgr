package colors

import (
	"fmt"
	"math"
	"time"

	"image/color"

	ciede2000 "github.com/mattn/go-ciede2000"
)

type Color256 struct {
	Color88
}

func (c *Color256) Supports(s uint8) *Color256 {
	c.Supported = s

	return c
}

func (c *Color256) NPrint(s ...string) string {

	if (c.fg == nil) && (c.bg == nil) {
		return s[0]
	}

	var closestColor256 func(*RGB, bool) int
	var debug bool = false

	closestColor256 = closestColor256Euclid
	if len(s) > 1 {
		if stringInSlice(s, "ciede2000") {
			closestColor256 = closestColor256CIEDE2000
		}
		if stringInSlice(s, "debug") {
			debug = true
		}
	}

	var r string

	if c.fg != nil {
		var fg int
		if debug {
			fg = closestColor256(c.fg, true)
		} else {
			fg = closestColor256(c.fg, false)
		}

		// ESC [ 3 8 ; 5 ; Ps m
		r = fmt.Sprintf("%s%s[38;5;%dm", r, escape, fg)
	}

	if c.bg != nil {
		var bg int
		if debug {
			bg = closestColor256(c.bg, true)
		} else {
			bg = closestColor256(c.bg, false)
		}

		// // ESC [ 4 8 ; 5 ; Ps m
		r = fmt.Sprintf("%s%s[48;5;%dm", r, escape, bg)
	}

	r = fmt.Sprintf("%s%s", r, s[0])

	return r
}

func (c *Color256) Print(s ...string) string {

	r := c.NPrint(s...)

	if c.bg != nil {
		// ESC [ 4 9 m
		r = fmt.Sprintf("%s%s[49m", r, escape)
	}

	if c.fg != nil {
		// ESC [ 3 9 m
		r = fmt.Sprintf("%s%s[39m", r, escape)
	}

	return r
}

func rgbDiffEuclid(a, b *RGB) float64 {
	return math.Sqrt(math.Pow((float64(a.r-b.r)*0.3), 2) + math.Pow((float64(a.g-b.g)*0.59), 2) + math.Pow((float64(a.b-b.b)*0.11), 2))
}

func getCode256(r, g, b uint8) uint8 {
	return (16 + (r * 36) + (g * 6) + b)
}

func getStep256(n uint8) uint8 {
	var r uint8
	if n < 95 {
		r = 0
	} else {
		r = (n - 55) / 40
	}

	if r == 5 {
		return 4
	}

	return r
}

func closestColor256Euclid(c *RGB, debug bool) int {

	/*
		TODO: optimize this to not run 256 * n times.
		      takes ~= 100 microsec
		var steps256 = []int{0, 95, 135, 175, 215, 25
		var r int
		sr := getStep256(c.r)
		sg := getStep256(c.r)
		sb := getStep256(c.
		c1 := &RGB{steps256[sr], steps256[sg], steps256[sb]}
		c2 := &RGB{steps256[sr+1], steps256[sg+1], steps256[sb+1
		d1 := rgbDiffEuclid(c, c1)
		d2 := rgbDiffEuclid(c, c
		dmin := math.Min(d1, d
		if d1 < d2 {
			r = 16 + (36 * sr) + (6 * sg) + sb
		} else {
			r = 16 + (36 * (sr + 1)) + (6 * (sg + 1)) + (sb + 1)
		}
	*/

	start := time.Now()

	dmin := math.Inf(1)
	r := -1
	for i := 16; i < 256; i++ {
		ia := Color256Lookup[i]
		cn := &RGB{ia[0], ia[1], ia[2]}
		dn := rgbDiffEuclid(c, cn)

		if dn < dmin {
			dmin = dn
			r = i
		}
	}

	if debug {
		elapsed := time.Since(start)
		fmt.Println(elapsed, c, Color256Lookup[r])
	}

	return r
}

func closestColor256CIEDE2000(c *RGB, debug bool) int {

	/*
		TODO: optimize this to not run 256 * n times.
		      takes ~= 1 millisec
	*/

	start := time.Now()

	dmin := math.Inf(1)
	r := -1
	c1 := &color.RGBA{uint8(c.r), uint8(c.g), uint8(c.b), 255}
	for i := 16; i < 256; i++ {
		ia := Color256Lookup[i]
		cn := &color.RGBA{uint8(ia[0]), uint8(ia[1]), uint8(ia[2]), 255}

		dn := ciede2000.Diff(c1, cn)

		if dn < dmin {
			dmin = dn
			r = i
		}
	}

	if debug {
		elapsed := time.Since(start)
		fmt.Println(elapsed, c, Color256Lookup[r])
	}

	return r
}
