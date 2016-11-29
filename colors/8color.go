package colors

import (
	"fmt"
	"image/color"
	"math"
	"strings"
	"time"

	ciede2000 "github.com/mattn/go-ciede2000"
)

type Color8 struct {
	fg, bg    *RGB
	Supported uint8
	color8    sysColor
}

var Color8Lookup map[string]*[4]uint8

func init() {
	Color8Lookup = make(map[string]*[4]uint8)

	Color8Lookup["black"] = &[4]uint8{30, 0, 0, 0}
	Color8Lookup["red"] = &[4]uint8{31, 205, 0, 0}
	Color8Lookup["green"] = &[4]uint8{32, 0, 205, 0}
	Color8Lookup["yellow"] = &[4]uint8{33, 205, 205, 0}
	Color8Lookup["blue"] = &[4]uint8{34, 0, 0, 238}
	Color8Lookup["magenta"] = &[4]uint8{35, 205, 0, 205}
	Color8Lookup["cyan"] = &[4]uint8{36, 0, 205, 205}
	Color8Lookup["grey"] = &[4]uint8{37, 229, 229, 229}
}

func (c *Color8) Supports(s uint8) *Color8 {
	c.Supported = s

	return c
}

func parseRGB8(name string) *[4]uint8 {
	rgb, exists := Color8Lookup[strings.ToLower(name)]
	if exists == false {
		return nil
	}
	return rgb
}

func (c *Color8) Fg8(s string) *Color8 {

	c.color8.fg = parseRGB8(s)

	return c
}

func (c *Color8) Bg8(s string) *Color8 {

	c.color8.bg = parseRGB8(s)

	return c
}

func (c *Color8) NPrint(s ...string) string {

	if (c.fg == nil) && (c.bg == nil) {
		return s[0]
	}

	var closestColor8 func(*RGB, bool) uint8
	var debug bool = false

	closestColor8 = closestColor8Euclid
	if len(s) > 1 {
		if stringInSlice(s, "ciede2000") {
			closestColor8 = closestColor8CIEDE2000
		}
		if stringInSlice(s, "debug") {
			debug = true
		}
	}

	var r string

	if c.color8.fg != nil { // use given fg
		// ESC [ Ps m
		r = fmt.Sprintf("%s%s[%dm", r, escape, c.color8.fg[0])
		if debug {
			fmt.Println(c.bg, c.color8.fg)
		}
	} else { // find closest fg
		if c.fg != nil {
			var fg uint8
			if debug {
				fg = closestColor8(c.fg, true)
			} else {
				fg = closestColor8(c.fg, false)
			}

			// ESC [ Ps m
			r = fmt.Sprintf("%s%s[%dm", r, escape, fg)
		} // c.fg != nil
	}

	if c.color8.bg != nil { // use given bg
		// ESC [ Ps m
		r = fmt.Sprintf("%s%s[%dm", r, escape, c.color8.bg[0]+10)
		if debug {
			fmt.Println(c.bg, c.color8.bg)
		}
	} else { // find closest bg
		if c.bg != nil {
			var bg uint8
			if debug {
				bg = closestColor8(c.bg, true)
			} else {
				bg = closestColor8(c.bg, false)
			}

			// ESC [ Ps m
			r = fmt.Sprintf("%s%s[%dm", r, escape, bg+10)
		} // c.fg != nil
	}

	r = fmt.Sprintf("%s%s", r, s[0])

	return r
}

func (c *Color8) Print(s ...string) string {
	r := c.NPrint(s...)

	if c.color8.bg != nil || c.bg != nil {
		// ESC [ 4 9 m
		r = fmt.Sprintf("%s%s[49m", r, escape)
	}

	if c.color8.fg != nil || c.fg != nil {
		// ESC [ 3 9 m
		r = fmt.Sprintf("%s%s[39m", r, escape)
	}

	return r
}

func closestColor8Euclid(c *RGB, debug bool) uint8 {
	start := time.Now()

	dmin := math.Inf(1)
	r := uint8(0)
	s := ""

	for k, v := range Color8Lookup {

		cn := &RGB{v[1], v[2], v[3]}
		dn := rgbDiffEuclid(c, cn)

		if dn < dmin {
			dmin = dn
			r = v[0]
			s = k
		}
	}

	if debug {
		elapsed := time.Since(start)
		fmt.Println(elapsed, c, Color8Lookup[s])
	}

	return r
}

func closestColor8CIEDE2000(c *RGB, debug bool) uint8 {
	start := time.Now()

	dmin := math.Inf(1)
	r := uint8(0)
	s := ""
	c1 := &color.RGBA{uint8(c.r), uint8(c.g), uint8(c.b), 255}

	for k, v := range Color8Lookup {

		cn := &color.RGBA{uint8(v[1]), uint8(v[2]), uint8(v[3]), 255}

		dn := ciede2000.Diff(c1, cn)

		if dn < dmin {
			dmin = dn
			r = v[0]
			s = k
		}
	}

	if debug {
		elapsed := time.Since(start)
		fmt.Println(elapsed, c, Color8Lookup[s])
	}

	return r
}
