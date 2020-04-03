#!/usr/bin/env bash

set -e
set -o pipefail

# For Linux ; on MacOS, use homebrew
# golangci-lint has version 1.24.0 built from 6fd4383 on 2020-03-15T11:38:02Z

if [[ -x bin/golangci-lint ]] ; then
    case "$(bin/golangci-lint --version)" in
    *1.24.0* ) exit 0 ;;
    esac
fi

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
