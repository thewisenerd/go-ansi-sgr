#! /bin/bash

X11URL="https://cgit.freedesktop.org/xorg/app/rgb/plain/rgb.txt"
W3URL="https://www.w3.org/TR/css3-color/"

X11F="rgb.txt"
W3F="css3-color.html"

TEMP="./tmp"

mkdir -p ${TEMP}

curl --silent ${X11URL} > ${TEMP}/${X11F}
curl --silent ${W3URL} > ${TEMP}/${W3F}
