#!/usr/bin/env bash

echo "Using GOPATH=$GOPATH"
echo "Getting github.com/inkeliz-technologies/tango using 'go get'"
go get -t -v ./... || exit 1

# These can fail without us minding it
blacklist="github.com/inkeliz-technologies/tango/demos/demoutils"

if [ "$TEST_TYPE" == "linux_test" ]
then
    echo "Testing github.com/inkeliz-technologies/tango using coveralls"
    $HOME/gopath/bin/goveralls -service=travis-ci

    echo "Testing and benchmarking github.com/inkeliz-technologies/tango"
    go test -v -bench=. ./... || exit 1

    echo "Checking for unnecessary conversions using unconvert"
    unconvert -v github.com/inkeliz-technologies/tango
elif [ "$TEST_TYPE" == "linux_build" ]
then
    for dir in `pwd`/demos/*/
    do
        # Formatting the directory to be usable by Go
        dir=${dir%*/}
        dir=${dir#$GOPATH/src/}

        # Ignore the directory if it's in the blacklist
        if [[ $blacklist == *"${dir}"* ]]
        then
            echo "Skipping ${dir}"
            continue
        fi

        # Some debug output and output directory initialization
        echo "Verifying ${dir} ..."
        outdir="/tmp/go-builds"

        # Creating the output directory, attempting to build and exit 1 if it failed
        mkdir -p "$outdir/linux/"
        go build -o "$outdir/linux/${dir}" -tags demo ${dir} || exit 1
    done
elif [ "$TEST_TYPE" == "js_test" ]
then
    echo "Getting and installing node.js"
    wget -qO- https://raw.githubusercontent.com/creationix/nvm/v0.34.0/install.sh | bash
    source ~/.nvm/nvm.sh
    nvm install 10
    npm install -g source-map-support
    GOOS=js GOARCH=wasm go test -v -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" --bench=. ./... || exit 1
elif [ "$TEST_TYPE" == "js_build" ]
then
    for dir in `pwd`/demos/*/
    do
        # Formatting the directory to be usable by Go
        dir=${dir%*/}
        dir=${dir#$GOPATH/src/}

        # Ignore the directory if it's in the blacklist
        if [[ $blacklist == *"${dir}"* ]]
        then
            echo "Skipping ${dir}"
            continue
        fi

        # Some debug output and output directory initialization
        echo "Verifying ${dir} ..."
        outdir="/tmp/go-builds"

        # Creating the output directory, attempting to build and exit 1 if it failed
        mkdir -p "$outdir/gopherjs/"
        GOOS=js GOARCH=wasm go build -o "$outdir/gopherjs/${dir}.wasm" --tags demo ${dir} || exit 1
    done
elif [ "$TEST_TYPE" == "android_test" ]
then
    echo "Skipping tests for github.com/inkeliz-technologies/tango using 'gomobile' (no tools exist yet)"
elif [ "$TEST_TYPE" == "android_build" ]
then
    for dir in `pwd`/demos/*/
    do
        # Formatting the directory to be usable by Go
        dir=${dir%*/}
        dir=${dir#$GOPATH/src/}

        # Ignore the directory if it's in the blacklist
        if [[ $blacklist == *"${dir}"* ]]
        then
            echo "Skipping ${dir}"
            continue
        fi

        # Some debug output and output directory initialization
        echo "Verifying ${dir} ..."
        outdir="/tmp/go-builds"

        # Creating the output directory, attempting to build and exit 1 if it failed
        mkdir -p `dirname "$outdir/android/${dir}.apk"`
        gomobile build -o "$outdir/android/${dir}.apk" -target android -tags demo ${dir} || exit 1
    done
elif [ "$TEST_TYPE" == "traffic_manager" ]
then
    branches='01-hello-world 02-first-system 03-camera-movement 04-hud 05-tilemaps 06-spritesheets-and-automated-city-building 07-hud-text'
    cd $HOME/gopath/src/github.com/inkeliz-technologies/TrafficManager
    for branch in $branches
    do
        echo "Verifying ${branch} ..."
        git checkout ${branch}
        go build -o "tmp/go-builds/${branch}" || exit 1
    done
else
    echo "environment variable TEST_TYPE was not set"
fi
