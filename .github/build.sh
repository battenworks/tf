#!/bin/bash
set -eux

PROJECT_NAME=$(basename $GITHUB_REPOSITORY)
NAME="${PROJECT_NAME}_${VERSION}_${GOOS}_${GOARCH}"

echo $GITHUB_REPOSITORY
echo $PROJECT_NAME

EXT=''

if [ $GOOS == 'windows' ]; then
  EXT='.exe'
fi

tar cvfz ${NAME}.tar.gz "${PROJECT_NAME}${EXT}" LICENSE
md5sum ${NAME}.tar.gz | cut -d ' ' -f 1 > ${NAME}_checksum.txt
