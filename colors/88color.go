package colors

import (
	"fmt"
	"math"
	"time"

	"image/color"

	ciede2000 "github.com/mattn/go-ciede2000"
)

type Color88 struct {
	Color16
}

func (c *Color88) Supports(s uint8) *Color88 {
	c.Supported = s

	return c
}

func (c *Color88) NPrint(s ...string) string {

	if (c.fg == nil) && (c.bg == nil) {
		return s[0]
	}

	var closestColor88 func(*RGB, bool) int
	var debug bool = false

	closestColor88 = closestColor88Euclid
	if len(s) > 1 {
		if stringInSlice(s, "ciede2000") {
			closestColor88 = closestColor88CIEDE2000
		}
		if stringInSlice(s, "debug") {
			debug = true
		}
	}

	var r string

	if c.fg != nil {
		var fg int
		if debug {
			fg = closestColor88(c.fg, true)
		} else {
			fg = closestColor88(c.fg, false)
		}

		// ESC [ 3 8 ; 5 ; Ps m
		r = fmt.Sprintf("%s%s[38;5;%dm", r, escape, fg)
	}

	if c.bg != nil {
		var bg int
		if debug {
			bg = closestColor88(c.bg, true)
		} else {
			bg = closestColor88(c.bg, false)
		}

		// ESC [ 4 8 ; 5 ; Ps m
		r = fmt.Sprintf("%s%s[48;5;%dm", r, escape, bg)
	}

	r = fmt.Sprintf("%s%s", r, s[0])

	return r
}

func (c *Color88) Print(s ...string) string {

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

func closestColor88Euclid(c *RGB, debug bool) int {

	/*
		TODO: optimize this to not run 256 * n times.
		      takes ~= 100 microsec
	*/

	start := time.Now()

	dmin := math.Inf(1)
	r := -1
	for i := 16; i < 88; i++ {
		ia := Color88Lookup[i]
		cn := &RGB{ia[0], ia[1], ia[2]}
		dn := rgbDiffEuclid(c, cn)

		if dn < dmin {
			dmin = dn
			r = i
		}
	}

	if debug {
		elapsed := time.Since(start)
		fmt.Println(elapsed, c, Color88Lookup[r])
	}

	return r
}

func closestColor88CIEDE2000(c *RGB, debug bool) int {

	/*
		TODO: optimize this to not run 256 * n times.
		      takes ~= 1 millisec
	*/

	start := time.Now()

	dmin := math.Inf(1)
	r := -1
	c1 := &color.RGBA{uint8(c.r), uint8(c.g), uint8(c.b), 255}
	for i := 16; i < 88; i++ {
		ia := Color88Lookup[i]
		cn := &color.RGBA{uint8(ia[0]), uint8(ia[1]), uint8(ia[2]), 255}

		dn := ciede2000.Diff(c1, cn)

		if dn < dmin {
			dmin = dn
			r = i
		}
	}

	if debug {
		elapsed := time.Since(start)
		fmt.Println(elapsed, c, Color88Lookup[r])
	}

	return r
}
