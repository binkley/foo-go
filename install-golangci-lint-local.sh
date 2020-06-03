#!/usr/bin/env bash

set -e
set -o pipefail

# For Linux ; on MacOS, use homebrew
# golangci-lint has version 1.27.0

if [[ -x bin/golangci-lint ]] ; then
    case "$(bin/golangci-lint --version)" in
    *1.27.0* ) exit 0 ;;
    esac
fi

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0
