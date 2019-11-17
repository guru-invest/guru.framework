#!/bin/bash

# Expects the directory structure:
# .
# └── projectname
#     ├── build.sh
#     └── src
#         ├── custompackage
#         │   └── custompackage.go
#         └── main
#             └── main.go
#
# This will build a binary called "projectname" at "projectname/bin/projectname".
#
# The final structure will look like:
#
# .
# └── projectname
#     ├── bin
#     │   └── projectname
#     ├── build.sh
#     └── src
#         ├── custompackage
#         │   └── custompackage.go
#         └── main
#             └── main.go
#

# Save the pwd before we run anything
PRE_PWD=`pwd`

# Determine the build script's actual directory, following symlinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

# Derive the project name from the directory
PROJECT="$(basename $BUILD_DIR)"

# Setup the environment for the build
GOPATH=$BUILD_DIR/vendor
export GOBIN=$BUILD_DIR/bin

# Build the project
cd $BUILD_DIR
mkdir -p bin

function build {
    cd $1
    BINARY="$(basename $1)"
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $GOBIN/$BINARY &>/dev/null
    cd $BUILD_DIR
    #go install $1 &>/dev/null
    EXIT_STATUS=$?
    if [ $EXIT_STATUS == 0 ]; then
      echo "Build succeeded"
    else
      getDirectory $1
    fi
}

function getDirectory {
    for f in $1/*; do
        if [ -d "$f" ]; then
            echo "Building $(basename $f)"
            go fmt
            build $f
        fi
    done
}

getDirectory $BUILD_DIR/src


# Change back to where we were
cd $PRE_PWD

exit 0