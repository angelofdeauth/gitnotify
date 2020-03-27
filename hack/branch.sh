#!/usr/bin/env bash
# @File:     branch.sh
# @Created:  2020-03-26 21:59:49
# @Modified: 2020-03-26 22:20:54
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

# Create a new branch from current.  Usage:
#
#   $ hack/branch.sh BRANCHNAME [ORIGIN]

git fetch
if git show-ref --verify --quiet refs/heads/"${1}"; then
  if [ -z "${2}" ]; then 
    git checkout -b "${1}" origin/"${1}"
    git branch --set-upstream "${1}" origin/"${1}"
  else
    git checkout -b "${1}" "${2}"/"${1}"
    git branch --set-upstream "${1}" "${2}"/"${1}"
  fi
else
  echo "[ERROR]: Branch ${1} already exists."
fi
