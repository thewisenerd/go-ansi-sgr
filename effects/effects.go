package effects

import "fmt"

const escape = "\x1b"

type fx [2]string

var effectLookup map[string]*fx

func init() {
	effectLookup = make(map[string]*fx)

	/*
	   Ps = 0  -> Normal (default).
	   Ps = 1  -> Bold.
	   Ps = 2  -> Faint, decreased intensity (ISO 6429).
	   Ps = 3  -> Italicized (ISO 6429).
	   Ps = 4  -> Underlined.
	   Ps = 5  -> Blink (appears as Bold).
	   Ps = 7  -> Inverse.
	   Ps = 8  -> Invisible, i.e., hidden (VT300).
	   Ps = 9  -> Crossed-out characters (ISO 6429).
	   Ps = 2 1  -> Doubly-underlined (ISO 6429).
	   Ps = 2 2  -> Normal (neither bold nor faint).
	   Ps = 2 3  -> Not italicized (ISO 6429).
	   Ps = 2 4  -> Not underlined.
	   Ps = 2 5  -> Steady (not blinking).
	   Ps = 2 7  -> Positive (not inverse).
	   Ps = 2 8  -> Visible, i.e., not hidden (VT300).
	   Ps = 2 9  -> Not crossed-out (ISO 6429).
	*/

	effectLookup["bold"] = &fx{"1", "22"}
	effectLookup["faint"] = &fx{"2", "22"}
	effectLookup["italic"] = &fx{"3", "23"}
	effectLookup["underline"] = &fx{"4", "24"}
	effectLookup["blink"] = &fx{"5", "25"}

	effectLookup["inverse"] = &fx{"7", "27"}
	effectLookup["hidden"] = &fx{"8", "28"}
	effectLookup["invisible"] = effectLookup["hidden"]
}

type Effect struct {
	params []*fx
}

func New(value ...string) *Effect {
	var e Effect

	for _, v := range value {
		effect, exists := effectLookup[v]
		if exists == false {
			continue
		}
		e.params = append(e.params, effect)
	}

	return &e
}

func (e *Effect) Add(value ...string) *Effect {
	for _, v := range value {
		effect, exists := effectLookup[v]
		if exists == false {
			continue
		}
		e.params = append(e.params, effect)
	}
	return e
}

func (e *Effect) Sprint(s string) string {
	// TODO: maketh Snprint

	var r string

	r = s

	// CSI Ps m ... CSI Ps m"
	for _, v := range e.params {
		r = fmt.Sprintf("%s[%sm%s%s[%sm", escape, v[0], r, escape, v[1])
	}
	return r
}

func (e *Effect) Reset() string {
	return Reset()
}

func Reset() string {
	return fmt.Sprintf("%s[0m", escape)
}
