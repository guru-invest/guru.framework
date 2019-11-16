#!/usr/bin/env bash

# Save the pwd before we run anything
PRE_PWD=`pwd`

# Determine the build script's actual directory, following symlinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

# Derive the project name from the directory
PROJECT="$(basename $BUILD_DIR)"

go mod init
go mod vendor

go get -u $1 >> /dev/null
mkdir -p $BUILD_DIR/vendor/$1
cp -r $GOPATH/src/$1/ $BUILD_DIR/vendor/$1