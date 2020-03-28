#!/usr/bin/env bash
# @File:     pr.sh
# @Created:  2020-03-27 01:04:32
# @Modified: 2020-03-27 20:36:27
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

set -ex

cd "$(dirname "$0")/../"

if [ -z "${BRANCH}" ]; then
  BRANCH=$(git rev-parse --abbrev-ref HEAD)
fi

if curl -XPOST -d "{\"title\":\"${BRANCH} automatic merge to master\", \"base\":\"master\", \"head\":\"${BRANCH}\", \"body\":\"${BRANCH} automatic merge to master\"}" \
  -s -o /dev/null -w "%{http_code}" https://api.github.com/repos/"$(git remote -v | grep origin | head -n1 | awk '{print $2}' | cut -d':' -f2 | rev | cut -c 5- | rev)"/pulls \
  -H "Authorization: token ${GITHUB_TOKEN}" ; then
  echo "[OK] PR created for branch ${BRANCH} successfully!"
else
  echo "[ERROR] PR creation for branch ${BRANCH} failed!"
fi
