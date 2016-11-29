#! /usr/bin/env bash

TEMP="./tmp"

X11F=${TEMP}/"rgb.txt"
W3F=${TEMP}/"css3-color.html"
TF=${TEMP}/"color-table.html"

cat ${TF} | grep -oP "<dfn id=[^>]*>\K([^<]*)" | while read -r color
do
  echo -en ${color}"|"
  rgb=$( grep -iP "\t${color}$" ${X11F} | tr -s ' ')
  echo $(echo $rgb | cut -d' ' -f1)", "$(echo $rgb | cut -d' ' -f2)", "$(echo $rgb | cut -d' ' -f3)
done
