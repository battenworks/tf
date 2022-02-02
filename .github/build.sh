#!/bin/bash
set -eux

echo $GITHUB_REPOSITORY
PROJECT_NAME=$(basename $GITHUB_REPOSITORY)
NAME="${PROJECT_NAME}_${VERSION}_${GOOS}_${GOARCH}"

EXT=''

if [ $GOOS == 'windows' ]; then
  EXT='.exe'
fi

tar cvfz ${NAME}.tar.gz "${PROJECT_NAME}${EXT}" LICENSE
md5sum ${NAME}.tar.gz | cut -d ' ' -f 1 > ${NAME}_checksum.txt
