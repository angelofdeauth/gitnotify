#!/usr/bin/env bash
# @File:     branch.sh
# @Created:  2020-03-26 21:59:49
# @Modified: 2020-03-31 16:38:37
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

# Create a new branch from current.  Usage:
#
#   $ hack/branch.sh BRANCHNAME

set -ex

cd "$(dirname "$0")/../"

# check if branch exists or is empty
if (! git show-ref --verify --quiet refs/heads/"${BRANCH}") && [ -n "${BRANCH}" ]; then
  
  # Create new branch, check it out, and set upstream
  git checkout -b "${BRANCH}"
  git push --set-upstream origin "${BRANCH}"

else

  # error out
  echo "[ERROR] Branch ${BRANCH} is null or already exists."

fi
