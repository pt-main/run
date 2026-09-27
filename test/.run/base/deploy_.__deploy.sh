#!/usr/bin/env bash
# A bash test script: reports args and its own name.
echo "BASH: script = $(basename "$0")"
echo "BASH: argc = $#"
i=1
for a in "$@"; do
  echo "BASH: arg[$i] = $a"
  i=$((i+1))
done
echo "BASH: pwd = $(pwd)"
