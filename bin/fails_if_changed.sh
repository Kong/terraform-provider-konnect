#!/bin/bash

if git diff --no-ext-diff --quiet -- .; then
  echo "No changes"
  exit
else
  echo "Checked in files are out of date from generated files, run ./tools/run.sh locally and check them in:"
  git status -s
  exit 1
fi