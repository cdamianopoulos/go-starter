#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

export GONOSUMDB="github.com/applatform/*"

# Create directory & set permissions.
mkdir -p ~/.ssh && chmod 0600 ~/.ssh

# Download the Github public SSH key.
ssh-keyscan -p 443 ssh.github.com >> ~/.ssh/known_hosts

# Force SSH to use port 443.
echo "Host github.com
  Hostname ssh.github.com
  Port 443
  User git" >> ~/.ssh/config

# Force git to use SSH.
git config --global url.ssh://git@github.com/.insteadOf https://github.com/