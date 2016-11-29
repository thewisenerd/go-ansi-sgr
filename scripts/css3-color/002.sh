#! /usr/bin/env bash

TEMP="./tmp"

X11F=${TEMP}/"rgb.txt"
W3F=${TEMP}/"css3-color.html"
TF=${TEMP}/"color-table.html"

S0=$( grep -nr "<h3.*Extended color keywords</h3>" ${W3F} | cut -d ':' -f1 )
S1=$( tail -n +${S0} ${W3F} | grep -n "<table class=colortable>" | cut -d':' -f1 )
L=$( tail -n +${S0} ${W3F} | grep -n "<h3.*4.4..*" | cut -d':' -f1 )

S=$( echo ${S0} + ${S1} - 1 | bc )
E=$( echo ${S0} + ${L} - 2 | bc )

sed -n -e ${S}','${E}'p' ${W3F} > ${TF}
