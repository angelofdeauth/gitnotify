#!/usr/bin/env bash
# @File:     dir-prep.sh
# @Created:  2020-03-25 17:25:25
# @Modified: 2020-03-26 19:28:00
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

set -ex

cd "$(dirname "$0")/../"

dirPrep() {
  mkdir -p "${BUILD_PATH}" && touch "${BUILD_PATH}"/make.log
  date -u
}

# Prepare build and log directories
if [ "$RELEASE" = "true" ]; then
  mkdir -p "${LOG_PATH}" && touch "${LOG_PATH}"/make.log 2>&1 | tee -a "${BUILD_PATH}"/make.log
  dirPrep 2>&1 | tee -a "${BUILD_PATH}"/make.log
else
  dirPrep 2>&1 | tee -a "${BUILD_PATH}"/make.log
fi
