#!/usr/bin/env bash
# @File:     pr.sh
# @Created:  2020-03-27 01:04:32
# @Modified: 2020-03-27 01:10:43
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

set -ex

cd "$(dirname "$0")/../"

if [ -z "${NAME}" ]; then
  NAME=$(git rev-parse --abbrev-ref HEAD)
  curl -XPOST -d "{\"title\":\"${NAME} automatic merge to master\", \"base\":\"master\", \"head\":\"${NAME}\", \"body\":\"${NAME} automatic merge to master\"}" \
    https://api.github.com/repos/"$(git remote -v | grep origin | head -n1 | awk '{print $2}' | cut -d':' -f2 | rev | cut -c 5- | rev)"/pulls \
    -H "Authorization: token ${GITHUB_TOKEN}"
elif [ -n "${NAME}" ]; then
  curl -XPOST -d "{\"title\":\"${NAME} automatic merge to master\", \"base\":\"master\", \"head\":\"${NAME}\", \"body\":\"${NAME} automatic merge to master\"}" \
    https://api.github.com/repos/"$(git remote -v | grep origin | head -n1 | awk '{print $2}' | cut -d':' -f2 | rev | cut -c 5- | rev)"/pulls \
    -H "Authorization: token ${GITHUB_TOKEN}"
fi
