#!/bin/bash
set -eux

VERSION=${GITHUB_REF#"refs/tags/"}
VERSION=${VERSION#"refs/heads/"}
ARCHIVE_NAME="${BINARY_NAME}_${VERSION}_${GOOS}_${GOARCH}"
EXT=''

if [ $GOOS == 'windows' ]; then
  EXT='.exe'
fi

go build -ldflags "-X main.version=${VERSION}" -o ${ARTIFACT_DIR}/${BINARY_NAME}${EXT}

cd ${ARTIFACT_DIR}
tar cvfz ${ARCHIVE_NAME}.tar.gz "${BINARY_NAME}${EXT}"
sha256sum ${ARCHIVE_NAME}.tar.gz > ${ARCHIVE_NAME}.sha256
