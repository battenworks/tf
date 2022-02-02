#!/bin/bash
set -eux

foo=${GITHUB_REF#"refs/tags/"}
echo $foo

cd $BINARY_DIRECTORY
NAME="${BINARY_NAME}_${VERSION}_${GOOS}_${GOARCH}"
EXT=''

if [ $GOOS == 'windows' ]; then
  EXT='.exe'
fi

tar cvfz ${NAME}.tar.gz "${BINARY_NAME}${EXT}"
md5sum ${NAME}.tar.gz | cut -d ' ' -f 1 > ${NAME}_checksum.txt
