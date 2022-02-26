#!/usr/bin/env sh

# see: https://k6.io/docs/using-k6/metrics/

set -ex

MY_PATH=$(dirname "$0")
MY_PATH=$(cd "$MY_PATH" && cd .. && pwd)
echo "running from... $MY_PATH"

type k6 >/dev/null 2>&1 || { echo >&2 "I require [k6] but it's not installed. Aborting."; exit 1; }
k6 run "$MY_PATH"/_loadtest/test.js