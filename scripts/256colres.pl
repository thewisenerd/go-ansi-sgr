#! /usr/bin/perl
# $XTermId: 256colres.pl,v 1.16 2007/06/08 23:58:37 tom Exp $
# -----------------------------------------------------------------------------
# this file is part of xterm
#
# Copyright 1999-2002,2007 by Thomas E. Dickey
#
#                         All Rights Reserved
#
# Permission is hereby granted, free of charge, to any person obtaining a
# copy of this software and associated documentation files (the
# "Software"), to deal in the Software without restriction, including
# without limitation the rights to use, copy, modify, merge, publish,
# distribute, sublicense, and/or sell copies of the Software, and to
# permit persons to whom the Software is furnished to do so, subject to
# the following conditions:
#
# The above copyright notice and this permission notice shall be included
# in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
# OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
# IN NO EVENT SHALL THE ABOVE LISTED COPYRIGHT HOLDER(S) BE LIABLE FOR ANY
# CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
# TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
# SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
#
# Except as contained in this notice, the name(s) of the above copyright
# holders shall not be used in advertising or otherwise to promote the
# sale, use or other dealings in this Software without prior written
# authorization.
# -----------------------------------------------------------------------------

# Construct a header file defining default resources for the 256-color model
# of xterm.  This is modeled after the 256colors2.pl script.

# use the resources for colors 0-15 - usually more-or-less a
# reproduction of the standard ANSI colors, but possibly more
# pleasing shades

use strict;

our ( $line );
our ( $red, $green, $blue, $gray );
our ( $level, $code, @steps );

print <<EOF;
/*
 * This header file was generated by $0
 */

package colors

var Color256Lookup []*[3]uint8

func init() {
	Color256Lookup = make([]*[3]uint8, 256)

EOF

$line="\tColor256Lookup[%d] = &[3]uint8{%d, %d, %d}\n";

print <<EOF;
	/* default system(xterm) colors. WILL vary from others */
	Color256Lookup[0] = &[3]uint8{0, 0, 0}			// black
	Color256Lookup[1] = &[3]uint8{205, 0, 0}			// red3
	Color256Lookup[2] = &[3]uint8{0, 205, 0}			// green3
	Color256Lookup[3] = &[3]uint8{205, 205, 0}		// yellow3
	Color256Lookup[4] = &[3]uint8{0, 0, 238}			// blue2
	Color256Lookup[5] = &[3]uint8{205, 0, 205}		// magenta3
	Color256Lookup[6] = &[3]uint8{0, 205, 205}		// cyan3
	Color256Lookup[7] = &[3]uint8{229, 229, 229}		// gray90
	Color256Lookup[8] = &[3]uint8{127, 127, 127}		// gray50
	Color256Lookup[9] = &[3]uint8{255, 0, 0}			// red
	Color256Lookup[10] = &[3]uint8{0, 255, 0}			// green
	Color256Lookup[11] = &[3]uint8{255, 255, 0}		// yellow
	Color256Lookup[12] = &[3]uint8{92, 92, 255}		// 5c/5c/ff
	Color256Lookup[13] = &[3]uint8{255, 255, 255}		// magenta
	Color256Lookup[14] = &[3]uint8{0, 255, 255}		// cyan
	Color256Lookup[15] = &[3]uint8{255, 255, 255}		// white
	/* default system(xterm) colors end.*/

EOF

# colors 16-231 are a 6x6x6 color cube
for ($red = 0; $red < 6; $red++) {
	for ($green = 0; $green < 6; $green++) {
		for ($blue = 0; $blue < 6; $blue++) {
			$code = 16 + ($red * 36) + ($green * 6) + $blue;
			printf($line,
				$code,
				($red ? ($red * 40 + 55) : 0),
				($green ? ($green * 40 + 55) : 0),
				($blue ? ($blue * 40 + 55) : 0)
			);
		}
	}
}

# colors 232-255 are a grayscale ramp, intentionally leaving out
# black and white
$code=232;
for ($gray = 0; $gray < 24; $gray++) {
	$level = ($gray * 10) + 8;
	$code = 232 + $gray;
	printf($line,
		$code, $level, $level, $level);
}

print <<EOF;
}
EOF