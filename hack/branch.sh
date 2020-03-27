#!/usr/bin/env bash
# @File:     branch.sh
# @Created:  2020-03-26 21:59:49
# @Modified: 2020-03-26 23:25:28
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

# Create a new branch from current.  Usage:
#
#   $ hack/branch.sh BRANCHNAME

set -ex

cd "$(dirname "$0")/../"

git pull
if ! git show-ref --verify --quiet refs/heads/"${1}"; then
  git checkout -b "${1}"
  git branch -u origin/"${1}"
else
  echo "[ERROR]: Branch ${1} already exists."
fi
