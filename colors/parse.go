package colors

import (
	"strconv"
	"strings"
)

func stringInSlice(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func parseRGB(name string) *RGB {
	rgb, exists := ColorCSSLookup[strings.ToLower(name)]
	if exists == false {
		return nil
	}
	return &RGB{rgb[0], rgb[1], rgb[2]}
}

func parseHex(hex string) *RGB {
	length := len(hex)

	if length == 0 {
		return nil
	}

	if hex[0] == '#' {
		hex = hex[1:]
		length--
	}

	if !((length == 3) || (length == 6)) {
		return nil
	}

	if length == 3 {
		r, e1 := strconv.ParseInt(string(hex[0]), 16, 8+1)
		g, e2 := strconv.ParseInt(string(hex[1]), 16, 8+1)
		b, e3 := strconv.ParseInt(string(hex[2]), 16, 8+1)

		if e1 != nil || e2 != nil || e3 != nil {
			return nil
		}

		return &RGB{
			16*uint8(r) + uint8(r),
			16*uint8(g) + uint8(g),
			16*uint8(b) + uint8(b),
		}
	} else if length == 6 {
		r, e1 := strconv.ParseInt(string(hex[0:2]), 16, 8+1)
		g, e2 := strconv.ParseInt(string(hex[2:4]), 16, 8+1)
		b, e3 := strconv.ParseInt(string(hex[4:6]), 16, 8+1)

		if e1 != nil || e2 != nil || e3 != nil {
			return nil
		}

		return &RGB{
			uint8(r),
			uint8(g),
			uint8(b),
		}
	}

	return nil
}
