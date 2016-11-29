package colors

import (
	"fmt"
)

func (c *Color) Supports(s uint8) *Color {
	c.Supported = s

	return c
}

func (c *Color) NPrint(s ...string) string {
	if (c.fg == nil) && (c.bg == nil) {
		return s[0]
	}

	var r string

	if c.fg != nil {
		// ESC [ 3 8; 2; Pr; Pg; Pb m
		r = fmt.Sprintf("%s%s[38;2;%d;%d;%dm", r, escape, c.fg.r, c.fg.g, c.fg.b)
	}

	if c.bg != nil {
		// ESC [ 4 8; 2; Pr; Pg; Pb m
		r = fmt.Sprintf("%s%s[48;2;%d;%d;%dm", r, escape, c.bg.r, c.bg.g, c.bg.b)
	}

	r = fmt.Sprintf("%s%s", r, s[0])

	return r
}

func (c *Color) Print(s ...string) string {

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
