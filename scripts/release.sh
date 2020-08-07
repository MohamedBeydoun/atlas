#!/bin/bash

VERSION=$1

if ! gr_loc="$(type -p "goreleaser")" || [[ -z $gr_loc ]]; then
    echo "Install goreleaser before running"
    exit 1
fi
if [ -z "$GITHUB_TOKEN" ]; then
    echo "Must set 'GITHUB_TOKEN'"
    exit 1
fi

git tag -a $VERSION -m "v$VERSION release"
git push origin $VERSION
ATLAS_VERSION=$VERSION goreleaser