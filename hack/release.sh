#!/usr/bin/env bash
# @File:     release.sh
# @Created:  2020-03-25 16:21:59
# @Modified: 2020-03-27 01:42:08
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

# Prepare for a release.  Usage:
#
#   $ hack/release.sh v0.1.0

set -ex

cd "$(dirname "$0")/../"

make build # ensure freshly-generated data
for GOOS in darwin linux windows; do
  GOARCH=amd64
  GOOS="${GOOS}" GOARCH="${GOARCH}" make build
  RELEASE_PATH="${BUILD_DIR}/${SEMVER}/${GOOS}_${GOARCH}"
  sha512sum "${RELEASE_PATH}"/* >"${RELEASE_PATH}"/release.sha512
  gpg --output "${RELEASE_PATH}"/release.sha512.sig --detach-sig "${RELEASE_PATH}"/release.sha512
done

git add .
git commit
git tag -sm "version ${SEMVER}" "${SEMVER}"

echo "[OK] Release build completed!"
