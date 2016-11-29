package colors

import (
	"fmt"
)

const escape = "\x1b"

type RGB struct {
	r, g, b uint8
}

type sysColor struct {
	fg, bg *[4]uint8
}

type Color struct {
	Color256
}

const (
	CTrue uint8 = iota
	C256
	C88
	C16
	C8
)

func (c *Color) Bg(name string) *Color {
	c.bg = parseRGB(name)

	return c
}

func (c *Color) Fg(name string) *Color {
	c.fg = parseRGB(name)

	return c
}

func (c *Color) BgHex(hex string) *Color {
	c.bg = parseHex(hex)

	return c
}

func (c *Color) FgHex(hex string) *Color {
	c.fg = parseHex(hex)

	return c
}

func (c *Color) Snprint(s ...string) string {

	if c.Supported == CTrue {
		return c.NPrint(s...)
	} else if c.Supported == C256 {
		return c.Color256.NPrint(s...)
	} else if c.Supported == C88 {
		return c.Color88.NPrint(s...)
	} else if c.Supported == C16 {
		return c.Color16.NPrint(s...)
	} else if c.Supported == C8 {
		return c.Color8.NPrint(s...)
	}

	// default to color256 if invalid c.Supported
	return c.Color256.NPrint(s...)
}

func (c *Color) Sprint(s ...string) string {

	if c.Supported == CTrue {
		return c.Print(s...)
	} else if c.Supported == C256 {
		return c.Color256.Print(s...)
	} else if c.Supported == C88 {
		return c.Color88.Print(s...)
	} else if c.Supported == C16 {
		return c.Color16.Print(s...)
	} else if c.Supported == C8 {
		return c.Color8.Print(s...)
	}

	// default to color256 if invalid c.Supported
	return c.Color256.Print(s...)
}

func (c *Color) Reset() string {
	return Reset()
}

func Reset() string {
	return fmt.Sprintf("%s[0m", escape)
}
