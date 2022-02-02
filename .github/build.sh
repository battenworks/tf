#!/bin/bash
set -eux

BINARY_NAME=$1
NAME="${BINARY_NAME}_${VERSION}_${GOOS}_${GOARCH}"

EXT=''

if [ $GOOS == 'windows' ]; then
  EXT='.exe'
fi

tar cvfz ${NAME}.tar.gz "${BINARY_NAME}${EXT}" LICENSE
md5sum ${NAME}.tar.gz | cut -d ' ' -f 1 > ${NAME}_checksum.txt
