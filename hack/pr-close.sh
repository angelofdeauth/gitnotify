#!/usr/bin/env bash
# @File:     pr-close.sh
# @Created:  2020-03-27 01:04:32
# @Modified: 2020-03-27 01:50:38
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

set -ex

cd "$(dirname "$0")/../"

if [ -z "${BRANCH}" ]; then
  BRANCH=$(git rev-parse --abbrev-ref HEAD)
fi

mergefn() {
  git fetch origin
  git checkout -b feature/build/make/create-prs-for-new-branches origin/feature/build/make/create-prs-for-new-branches
  git merge master
  git checkout master
  git merge --no-ff feature/build/make/create-prs-for-new-branches
  git push origin master
}

if mergefn; then
  echo "[OK] PR merged successfully!"
else
  echo "[ERROR] PR merge failed!"
fi
