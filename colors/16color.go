package colors

import (
	"fmt"
	"image/color"
	"math"
	"strings"
	"time"

	ciede2000 "github.com/mattn/go-ciede2000"
)

type Color16 struct {
	Color8
	color16 sysColor
}

var Color16Lookup map[string]*[4]uint8

func init() {
	Color16Lookup = make(map[string]*[4]uint8)

	Color16Lookup["black"] = &[4]uint8{30, 0, 0, 0}
	Color16Lookup["red"] = &[4]uint8{31, 205, 0, 0}
	Color16Lookup["green"] = &[4]uint8{32, 0, 205, 0}
	Color16Lookup["yellow"] = &[4]uint8{33, 205, 205, 0}
	Color16Lookup["blue"] = &[4]uint8{34, 0, 0, 238}
	Color16Lookup["magenta"] = &[4]uint8{35, 205, 0, 205}
	Color16Lookup["cyan"] = &[4]uint8{36, 0, 205, 205}
	Color16Lookup["grey"] = &[4]uint8{37, 229, 229, 229}

	Color16Lookup["darkgrey"] = &[4]uint8{90, 0, 0, 0}
	Color16Lookup["lightred"] = &[4]uint8{91, 205, 0, 0}
	Color16Lookup["lightgreen"] = &[4]uint8{92, 0, 205, 0}
	Color16Lookup["lightyellow"] = &[4]uint8{93, 205, 205, 0}
	Color16Lookup["lightblue"] = &[4]uint8{94, 0, 0, 238}
	Color16Lookup["lightmagenta"] = &[4]uint8{95, 205, 0, 205}
	Color16Lookup["lightcyan"] = &[4]uint8{96, 0, 205, 205}
	Color16Lookup["white"] = &[4]uint8{97, 229, 229, 229}

}

func (c *Color16) Supports(s uint8) *Color16 {
	c.Supported = s

	return c
}

func parseRGB16(name string) *[4]uint8 {
	rgb, exists := Color16Lookup[strings.ToLower(name)]
	if exists == false {
		return nil
	}
	return rgb
}

func (c *Color16) Fg16(s string) *Color16 {

	c.color16.fg = parseRGB16(s)

	return c
}

func (c *Color16) Bg16(s string) *Color16 {

	c.color16.bg = parseRGB16(s)

	return c
}

func (c *Color16) NPrint(s ...string) string {

	if (c.fg == nil) && (c.bg == nil) {
		return s[0]
	}

	var closestColor16 func(*RGB, bool) uint8
	var debug bool = false

	closestColor16 = closestColor16Euclid
	if len(s) > 1 {
		if stringInSlice(s, "ciede2000") {
			closestColor16 = closestColor16CIEDE2000
		}
		if stringInSlice(s, "debug") {
			debug = true
		}
	}

	var r string

	if c.color16.fg != nil { // use given fg
		// ESC [ Ps m
		r = fmt.Sprintf("%s%s[%dm", r, escape, c.color16.fg[0])
		if debug {
			fmt.Println(c.fg, c.color16.fg)
		}
	} else { // find closest fg
		if c.fg != nil {
			var fg uint8
			if debug {
				fg = closestColor16(c.fg, true)
			} else {
				fg = closestColor16(c.fg, false)
			}

			// ESC [ Ps m
			r = fmt.Sprintf("%s%s[%dm", r, escape, fg)
		} // c.fg != nil
	}

	if c.color16.bg != nil { // use given bg
		// ESC [ Ps m
		r = fmt.Sprintf("%s%s[%dm", r, escape, c.color16.bg[0]+10)
		if debug {
			fmt.Println(c.bg, c.color16.bg)
		}
	} else { // find closest bg
		if c.bg != nil {
			var bg uint8
			if debug {
				bg = closestColor16(c.bg, true)
			} else {
				bg = closestColor16(c.bg, false)
			}

			// ESC [ Ps m
			r = fmt.Sprintf("%s%s[%dm", r, escape, bg+10)
		} // c.fg != nil
	}

	r = fmt.Sprintf("%s%s", r, s[0])

	return r
}

func (c *Color16) Print(s ...string) string {
	r := c.NPrint(s...)

	if c.color16.bg != nil || c.bg != nil {
		// ESC [ 4 9 m
		r = fmt.Sprintf("%s%s[49m", r, escape)
	}

	if c.color16.fg != nil || c.fg != nil {
		// ESC [ 3 9 m
		r = fmt.Sprintf("%s%s[39m", r, escape)
	}

	return r
}

func closestColor16Euclid(c *RGB, debug bool) uint8 {
	start := time.Now()

	dmin := math.Inf(1)
	r := uint8(0)
	s := ""

	for k, v := range Color16Lookup {

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

func closestColor16CIEDE2000(c *RGB, debug bool) uint8 {
	start := time.Now()

	dmin := math.Inf(1)
	r := uint8(0)
	s := ""
	c1 := &color.RGBA{uint8(c.r), uint8(c.g), uint8(c.b), 255}

	for k, v := range Color16Lookup {

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
