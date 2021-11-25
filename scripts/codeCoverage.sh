#!/bin/bash -e
set -o errexit
set -o nounset
set -o pipefail

# Temporary filename to store the code coverage profile in for Go tests.
profile=c.out

# Execute all Go tests within the project.
 go test -race -cover -coverprofile=$profile ./...

# Open a browser window to show the code coverage result.
 go tool cover -html=$profile
