#!/usr/bin/env bash
# @File:     branch.sh
# @Created:  2020-03-26 21:59:49
# @Modified: 2020-03-27 19:31:52
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

# Create a new branch from current.  Usage:
#
#   $ hack/branch.sh BRANCHNAME

set -ex

cd "$(dirname "$0")/../"

if (! git show-ref --verify --quiet refs/heads/"${BRANCH}") && [ -n "${BRANCH}" ]; then
  git checkout -b "${BRANCH}"
  git push --set-upstream origin "${BRANCH}"
else
  echo "[ERROR] Branch ${BRANCH} is null or already exists."
fi
