#!/usr/bin/env bash
# @File:     commit.sh
# @Created:  2020-03-26 20:27:09
# @Modified: 2020-03-27 19:52:50
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

set  -ex

cd "$(dirname "$0")/../"

conditionalAddAll() {
  git status
  read -r -p "Add all files (y/n)?" choice
  case "$choice" in 
    y|Y ) echo "[OK] Adding all files..." && git add . && git commit && git push;;
    n|N ) echo "[ERROR] Exiting..."; exit 2;;
    * ) echo "[ERROR] Invalid, exiting..."; exit 3;;
  esac
}

# test if there are changed tracked files which are cached, 1 is cached changes and 0 is none.
if git diff --cached --quiet ; then
  
  # test if there are changed tracked files which are not cached, 1 is uncached changes and 0 is none.
  if git diff-index "$(git write-tree)" --quiet -- ; then
    if [ -n "$(git status -s)" ]; then
      echo "No changes detected either cached, new, or staged."
    else
      conditionalAddAll
    fi
  else
    conditionalAddAll
  fi
else
  git commit
  git push
fi

echo "[OK] Commit completed!"
