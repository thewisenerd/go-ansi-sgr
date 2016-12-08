package main

import (
	"fmt"

	"github.com/thewisenerd/go-ansi-sgr/colors"
	"github.com/thewisenerd/go-ansi-sgr/effects"
)

func main() {
	// TODO: maketh cli

	d := new(colors.Color).Fg("greenyellow").Bg("royalblue")
	d.Fg16("grey").Bg16("lightblue") // set 16color overrides (optional)
	d.Fg8("blue").Bg8("red")         // set 8color overrides (optional)

	d.Supports(colors.CTrue) // default
	fmt.Println(d.Sprint("truecolor"))

	d.Supports(colors.C256)                                   // 256-color mode
	fmt.Println(d.Sprint("256-color"))                        // default euclid
	fmt.Println(d.Sprint("256-color-ciede2000", "ciede2000", "debug")) // optional CIEDE2000 (more _expensive_)

	d.Supports(colors.C88)                                   // 88-color mode (will not work properly, unless actually 88-color)
	fmt.Println(d.Sprint("88-color"))                        // default euclid
	fmt.Println(d.Sprint("88-color-ciede2000", "ciede2000", "debug")) // optional CIEDE2000 (more _expensive_)

	d.Supports(colors.C16)                                   // 16-color mode
	fmt.Println(d.Sprint("16-color"))                        // default euclid (used if overrides not set)
	fmt.Println(d.Sprint("16-color-ciede2000", "ciede2000", "debug")) // optional CIEDE2000 (more _expensive_) (used if overrides not set)

	d.Supports(colors.C8)                                   // 8-color mode
	fmt.Println(d.Sprint("8-color"))                        // default euclid (used if overrides not set)
	fmt.Println(d.Sprint("8-color-ciede2000", "ciede2000", "debug")) // optional CIEDE2000 (more _expensive_) (used if overrides not set)

	fmt.Println(d.Color256.Print("256-color-direct"))                        // you can directly access the Print function of a color mode
	fmt.Println(d.Color256.Print("256-color-direct-ciede2000", "ciede2000", "debug")) // color difference algorithm can be passed too

	fmt.Println("---")

	d.Supports(colors.CTrue) // demo in true color glory

	fmt.Println(d.Snprint("n-print"))      // snprint does not add reset codes
	fmt.Println("carries color")           // so normal console out is affected
	fmt.Println(d.Snprint("across lines")) // unless color is reset

	fmt.Println("---") // {colors,effect}.Reset() outputs "ESC [ 0 m" which resets terminal to default colors

	fmt.Println(d.Snprint("n-print and reset"), d.Reset()) // snprint does not add reset codes
	fmt.Println("does not carries color")                  // so normal console out is affected
	fmt.Println(d.Snprint("across lines"), colors.Reset()) // unless color is reset
	fmt.Println("see ?")                                   // see ?

	fmt.Println("---")

	fmt.Println(effects.New("bold").Sprint("bold"))           // bold
	fmt.Println(effects.New("faint").Sprint("faint"))         // faint
	fmt.Println(effects.New("italic").Sprint("italic"))       // italic
	fmt.Println(effects.New("underline").Sprint("underline")) // underline
	fmt.Println(effects.New("blink").Sprint("blink"))         // blink

	fmt.Println(effects.New("inverse").Sprint("inverse"))                    // inverse
	fmt.Println(effects.New("hidden").Sprint("hidden"), "(hidden)")          // hidden
	fmt.Println(effects.New("invisible").Sprint("invisible"), "(invisible)") // invisible = hidden

	fmt.Println(effects.New("bold", "italic").Sprint("bold+italic")) // you can combine effects

	fx := new(effects.Effect) // you can _add_ effects to existing effects
	fmt.Println(fx.Sprint("noeffect"))
	fx.Add("bold", "underline" /*, ... */)
	fmt.Println(fx.Sprint("bold+underline"))

	fmt.Println("---")

	fmt.Println(d.Sprint(effects.New("bold", "italic").Sprint("bold+italic+color")))
	fmt.Println(d.Sprint(fx.Sprint("bold+underline+color")))

	fmt.Println("-end-")

}
