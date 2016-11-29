/*
 * This header file was generated by scripts/256colres.pl
 */

package colors

var Color256Lookup []*[3]uint8

func init() {
	Color256Lookup = make([]*[3]uint8, 256)

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

	Color256Lookup[16] = &[3]uint8{0, 0, 0}
	Color256Lookup[17] = &[3]uint8{0, 0, 95}
	Color256Lookup[18] = &[3]uint8{0, 0, 135}
	Color256Lookup[19] = &[3]uint8{0, 0, 175}
	Color256Lookup[20] = &[3]uint8{0, 0, 215}
	Color256Lookup[21] = &[3]uint8{0, 0, 255}
	Color256Lookup[22] = &[3]uint8{0, 95, 0}
	Color256Lookup[23] = &[3]uint8{0, 95, 95}
	Color256Lookup[24] = &[3]uint8{0, 95, 135}
	Color256Lookup[25] = &[3]uint8{0, 95, 175}
	Color256Lookup[26] = &[3]uint8{0, 95, 215}
	Color256Lookup[27] = &[3]uint8{0, 95, 255}
	Color256Lookup[28] = &[3]uint8{0, 135, 0}
	Color256Lookup[29] = &[3]uint8{0, 135, 95}
	Color256Lookup[30] = &[3]uint8{0, 135, 135}
	Color256Lookup[31] = &[3]uint8{0, 135, 175}
	Color256Lookup[32] = &[3]uint8{0, 135, 215}
	Color256Lookup[33] = &[3]uint8{0, 135, 255}
	Color256Lookup[34] = &[3]uint8{0, 175, 0}
	Color256Lookup[35] = &[3]uint8{0, 175, 95}
	Color256Lookup[36] = &[3]uint8{0, 175, 135}
	Color256Lookup[37] = &[3]uint8{0, 175, 175}
	Color256Lookup[38] = &[3]uint8{0, 175, 215}
	Color256Lookup[39] = &[3]uint8{0, 175, 255}
	Color256Lookup[40] = &[3]uint8{0, 215, 0}
	Color256Lookup[41] = &[3]uint8{0, 215, 95}
	Color256Lookup[42] = &[3]uint8{0, 215, 135}
	Color256Lookup[43] = &[3]uint8{0, 215, 175}
	Color256Lookup[44] = &[3]uint8{0, 215, 215}
	Color256Lookup[45] = &[3]uint8{0, 215, 255}
	Color256Lookup[46] = &[3]uint8{0, 255, 0}
	Color256Lookup[47] = &[3]uint8{0, 255, 95}
	Color256Lookup[48] = &[3]uint8{0, 255, 135}
	Color256Lookup[49] = &[3]uint8{0, 255, 175}
	Color256Lookup[50] = &[3]uint8{0, 255, 215}
	Color256Lookup[51] = &[3]uint8{0, 255, 255}
	Color256Lookup[52] = &[3]uint8{95, 0, 0}
	Color256Lookup[53] = &[3]uint8{95, 0, 95}
	Color256Lookup[54] = &[3]uint8{95, 0, 135}
	Color256Lookup[55] = &[3]uint8{95, 0, 175}
	Color256Lookup[56] = &[3]uint8{95, 0, 215}
	Color256Lookup[57] = &[3]uint8{95, 0, 255}
	Color256Lookup[58] = &[3]uint8{95, 95, 0}
	Color256Lookup[59] = &[3]uint8{95, 95, 95}
	Color256Lookup[60] = &[3]uint8{95, 95, 135}
	Color256Lookup[61] = &[3]uint8{95, 95, 175}
	Color256Lookup[62] = &[3]uint8{95, 95, 215}
	Color256Lookup[63] = &[3]uint8{95, 95, 255}
	Color256Lookup[64] = &[3]uint8{95, 135, 0}
	Color256Lookup[65] = &[3]uint8{95, 135, 95}
	Color256Lookup[66] = &[3]uint8{95, 135, 135}
	Color256Lookup[67] = &[3]uint8{95, 135, 175}
	Color256Lookup[68] = &[3]uint8{95, 135, 215}
	Color256Lookup[69] = &[3]uint8{95, 135, 255}
	Color256Lookup[70] = &[3]uint8{95, 175, 0}
	Color256Lookup[71] = &[3]uint8{95, 175, 95}
	Color256Lookup[72] = &[3]uint8{95, 175, 135}
	Color256Lookup[73] = &[3]uint8{95, 175, 175}
	Color256Lookup[74] = &[3]uint8{95, 175, 215}
	Color256Lookup[75] = &[3]uint8{95, 175, 255}
	Color256Lookup[76] = &[3]uint8{95, 215, 0}
	Color256Lookup[77] = &[3]uint8{95, 215, 95}
	Color256Lookup[78] = &[3]uint8{95, 215, 135}
	Color256Lookup[79] = &[3]uint8{95, 215, 175}
	Color256Lookup[80] = &[3]uint8{95, 215, 215}
	Color256Lookup[81] = &[3]uint8{95, 215, 255}
	Color256Lookup[82] = &[3]uint8{95, 255, 0}
	Color256Lookup[83] = &[3]uint8{95, 255, 95}
	Color256Lookup[84] = &[3]uint8{95, 255, 135}
	Color256Lookup[85] = &[3]uint8{95, 255, 175}
	Color256Lookup[86] = &[3]uint8{95, 255, 215}
	Color256Lookup[87] = &[3]uint8{95, 255, 255}
	Color256Lookup[88] = &[3]uint8{135, 0, 0}
	Color256Lookup[89] = &[3]uint8{135, 0, 95}
	Color256Lookup[90] = &[3]uint8{135, 0, 135}
	Color256Lookup[91] = &[3]uint8{135, 0, 175}
	Color256Lookup[92] = &[3]uint8{135, 0, 215}
	Color256Lookup[93] = &[3]uint8{135, 0, 255}
	Color256Lookup[94] = &[3]uint8{135, 95, 0}
	Color256Lookup[95] = &[3]uint8{135, 95, 95}
	Color256Lookup[96] = &[3]uint8{135, 95, 135}
	Color256Lookup[97] = &[3]uint8{135, 95, 175}
	Color256Lookup[98] = &[3]uint8{135, 95, 215}
	Color256Lookup[99] = &[3]uint8{135, 95, 255}
	Color256Lookup[100] = &[3]uint8{135, 135, 0}
	Color256Lookup[101] = &[3]uint8{135, 135, 95}
	Color256Lookup[102] = &[3]uint8{135, 135, 135}
	Color256Lookup[103] = &[3]uint8{135, 135, 175}
	Color256Lookup[104] = &[3]uint8{135, 135, 215}
	Color256Lookup[105] = &[3]uint8{135, 135, 255}
	Color256Lookup[106] = &[3]uint8{135, 175, 0}
	Color256Lookup[107] = &[3]uint8{135, 175, 95}
	Color256Lookup[108] = &[3]uint8{135, 175, 135}
	Color256Lookup[109] = &[3]uint8{135, 175, 175}
	Color256Lookup[110] = &[3]uint8{135, 175, 215}
	Color256Lookup[111] = &[3]uint8{135, 175, 255}
	Color256Lookup[112] = &[3]uint8{135, 215, 0}
	Color256Lookup[113] = &[3]uint8{135, 215, 95}
	Color256Lookup[114] = &[3]uint8{135, 215, 135}
	Color256Lookup[115] = &[3]uint8{135, 215, 175}
	Color256Lookup[116] = &[3]uint8{135, 215, 215}
	Color256Lookup[117] = &[3]uint8{135, 215, 255}
	Color256Lookup[118] = &[3]uint8{135, 255, 0}
	Color256Lookup[119] = &[3]uint8{135, 255, 95}
	Color256Lookup[120] = &[3]uint8{135, 255, 135}
	Color256Lookup[121] = &[3]uint8{135, 255, 175}
	Color256Lookup[122] = &[3]uint8{135, 255, 215}
	Color256Lookup[123] = &[3]uint8{135, 255, 255}
	Color256Lookup[124] = &[3]uint8{175, 0, 0}
	Color256Lookup[125] = &[3]uint8{175, 0, 95}
	Color256Lookup[126] = &[3]uint8{175, 0, 135}
	Color256Lookup[127] = &[3]uint8{175, 0, 175}
	Color256Lookup[128] = &[3]uint8{175, 0, 215}
	Color256Lookup[129] = &[3]uint8{175, 0, 255}
	Color256Lookup[130] = &[3]uint8{175, 95, 0}
	Color256Lookup[131] = &[3]uint8{175, 95, 95}
	Color256Lookup[132] = &[3]uint8{175, 95, 135}
	Color256Lookup[133] = &[3]uint8{175, 95, 175}
	Color256Lookup[134] = &[3]uint8{175, 95, 215}
	Color256Lookup[135] = &[3]uint8{175, 95, 255}
	Color256Lookup[136] = &[3]uint8{175, 135, 0}
	Color256Lookup[137] = &[3]uint8{175, 135, 95}
	Color256Lookup[138] = &[3]uint8{175, 135, 135}
	Color256Lookup[139] = &[3]uint8{175, 135, 175}
	Color256Lookup[140] = &[3]uint8{175, 135, 215}
	Color256Lookup[141] = &[3]uint8{175, 135, 255}
	Color256Lookup[142] = &[3]uint8{175, 175, 0}
	Color256Lookup[143] = &[3]uint8{175, 175, 95}
	Color256Lookup[144] = &[3]uint8{175, 175, 135}
	Color256Lookup[145] = &[3]uint8{175, 175, 175}
	Color256Lookup[146] = &[3]uint8{175, 175, 215}
	Color256Lookup[147] = &[3]uint8{175, 175, 255}
	Color256Lookup[148] = &[3]uint8{175, 215, 0}
	Color256Lookup[149] = &[3]uint8{175, 215, 95}
	Color256Lookup[150] = &[3]uint8{175, 215, 135}
	Color256Lookup[151] = &[3]uint8{175, 215, 175}
	Color256Lookup[152] = &[3]uint8{175, 215, 215}
	Color256Lookup[153] = &[3]uint8{175, 215, 255}
	Color256Lookup[154] = &[3]uint8{175, 255, 0}
	Color256Lookup[155] = &[3]uint8{175, 255, 95}
	Color256Lookup[156] = &[3]uint8{175, 255, 135}
	Color256Lookup[157] = &[3]uint8{175, 255, 175}
	Color256Lookup[158] = &[3]uint8{175, 255, 215}
	Color256Lookup[159] = &[3]uint8{175, 255, 255}
	Color256Lookup[160] = &[3]uint8{215, 0, 0}
	Color256Lookup[161] = &[3]uint8{215, 0, 95}
	Color256Lookup[162] = &[3]uint8{215, 0, 135}
	Color256Lookup[163] = &[3]uint8{215, 0, 175}
	Color256Lookup[164] = &[3]uint8{215, 0, 215}
	Color256Lookup[165] = &[3]uint8{215, 0, 255}
	Color256Lookup[166] = &[3]uint8{215, 95, 0}
	Color256Lookup[167] = &[3]uint8{215, 95, 95}
	Color256Lookup[168] = &[3]uint8{215, 95, 135}
	Color256Lookup[169] = &[3]uint8{215, 95, 175}
	Color256Lookup[170] = &[3]uint8{215, 95, 215}
	Color256Lookup[171] = &[3]uint8{215, 95, 255}
	Color256Lookup[172] = &[3]uint8{215, 135, 0}
	Color256Lookup[173] = &[3]uint8{215, 135, 95}
	Color256Lookup[174] = &[3]uint8{215, 135, 135}
	Color256Lookup[175] = &[3]uint8{215, 135, 175}
	Color256Lookup[176] = &[3]uint8{215, 135, 215}
	Color256Lookup[177] = &[3]uint8{215, 135, 255}
	Color256Lookup[178] = &[3]uint8{215, 175, 0}
	Color256Lookup[179] = &[3]uint8{215, 175, 95}
	Color256Lookup[180] = &[3]uint8{215, 175, 135}
	Color256Lookup[181] = &[3]uint8{215, 175, 175}
	Color256Lookup[182] = &[3]uint8{215, 175, 215}
	Color256Lookup[183] = &[3]uint8{215, 175, 255}
	Color256Lookup[184] = &[3]uint8{215, 215, 0}
	Color256Lookup[185] = &[3]uint8{215, 215, 95}
	Color256Lookup[186] = &[3]uint8{215, 215, 135}
	Color256Lookup[187] = &[3]uint8{215, 215, 175}
	Color256Lookup[188] = &[3]uint8{215, 215, 215}
	Color256Lookup[189] = &[3]uint8{215, 215, 255}
	Color256Lookup[190] = &[3]uint8{215, 255, 0}
	Color256Lookup[191] = &[3]uint8{215, 255, 95}
	Color256Lookup[192] = &[3]uint8{215, 255, 135}
	Color256Lookup[193] = &[3]uint8{215, 255, 175}
	Color256Lookup[194] = &[3]uint8{215, 255, 215}
	Color256Lookup[195] = &[3]uint8{215, 255, 255}
	Color256Lookup[196] = &[3]uint8{255, 0, 0}
	Color256Lookup[197] = &[3]uint8{255, 0, 95}
	Color256Lookup[198] = &[3]uint8{255, 0, 135}
	Color256Lookup[199] = &[3]uint8{255, 0, 175}
	Color256Lookup[200] = &[3]uint8{255, 0, 215}
	Color256Lookup[201] = &[3]uint8{255, 0, 255}
	Color256Lookup[202] = &[3]uint8{255, 95, 0}
	Color256Lookup[203] = &[3]uint8{255, 95, 95}
	Color256Lookup[204] = &[3]uint8{255, 95, 135}
	Color256Lookup[205] = &[3]uint8{255, 95, 175}
	Color256Lookup[206] = &[3]uint8{255, 95, 215}
	Color256Lookup[207] = &[3]uint8{255, 95, 255}
	Color256Lookup[208] = &[3]uint8{255, 135, 0}
	Color256Lookup[209] = &[3]uint8{255, 135, 95}
	Color256Lookup[210] = &[3]uint8{255, 135, 135}
	Color256Lookup[211] = &[3]uint8{255, 135, 175}
	Color256Lookup[212] = &[3]uint8{255, 135, 215}
	Color256Lookup[213] = &[3]uint8{255, 135, 255}
	Color256Lookup[214] = &[3]uint8{255, 175, 0}
	Color256Lookup[215] = &[3]uint8{255, 175, 95}
	Color256Lookup[216] = &[3]uint8{255, 175, 135}
	Color256Lookup[217] = &[3]uint8{255, 175, 175}
	Color256Lookup[218] = &[3]uint8{255, 175, 215}
	Color256Lookup[219] = &[3]uint8{255, 175, 255}
	Color256Lookup[220] = &[3]uint8{255, 215, 0}
	Color256Lookup[221] = &[3]uint8{255, 215, 95}
	Color256Lookup[222] = &[3]uint8{255, 215, 135}
	Color256Lookup[223] = &[3]uint8{255, 215, 175}
	Color256Lookup[224] = &[3]uint8{255, 215, 215}
	Color256Lookup[225] = &[3]uint8{255, 215, 255}
	Color256Lookup[226] = &[3]uint8{255, 255, 0}
	Color256Lookup[227] = &[3]uint8{255, 255, 95}
	Color256Lookup[228] = &[3]uint8{255, 255, 135}
	Color256Lookup[229] = &[3]uint8{255, 255, 175}
	Color256Lookup[230] = &[3]uint8{255, 255, 215}
	Color256Lookup[231] = &[3]uint8{255, 255, 255}
	Color256Lookup[232] = &[3]uint8{8, 8, 8}
	Color256Lookup[233] = &[3]uint8{18, 18, 18}
	Color256Lookup[234] = &[3]uint8{28, 28, 28}
	Color256Lookup[235] = &[3]uint8{38, 38, 38}
	Color256Lookup[236] = &[3]uint8{48, 48, 48}
	Color256Lookup[237] = &[3]uint8{58, 58, 58}
	Color256Lookup[238] = &[3]uint8{68, 68, 68}
	Color256Lookup[239] = &[3]uint8{78, 78, 78}
	Color256Lookup[240] = &[3]uint8{88, 88, 88}
	Color256Lookup[241] = &[3]uint8{98, 98, 98}
	Color256Lookup[242] = &[3]uint8{108, 108, 108}
	Color256Lookup[243] = &[3]uint8{118, 118, 118}
	Color256Lookup[244] = &[3]uint8{128, 128, 128}
	Color256Lookup[245] = &[3]uint8{138, 138, 138}
	Color256Lookup[246] = &[3]uint8{148, 148, 148}
	Color256Lookup[247] = &[3]uint8{158, 158, 158}
	Color256Lookup[248] = &[3]uint8{168, 168, 168}
	Color256Lookup[249] = &[3]uint8{178, 178, 178}
	Color256Lookup[250] = &[3]uint8{188, 188, 188}
	Color256Lookup[251] = &[3]uint8{198, 198, 198}
	Color256Lookup[252] = &[3]uint8{208, 208, 208}
	Color256Lookup[253] = &[3]uint8{218, 218, 218}
	Color256Lookup[254] = &[3]uint8{228, 228, 228}
	Color256Lookup[255] = &[3]uint8{238, 238, 238}
}
