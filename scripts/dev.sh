#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

docker-compose run --rm --entrypoint "" --service-ports dev bash
