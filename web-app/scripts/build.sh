#!/bin/bash

script_dir="${0%/*}"
cd $script_dir && \
    echo "Installing dependencies, this might take a few minutes..." && \
    cd ./.. && npm install && \

    echo "Building web app..." && \
    ng build -prod && \
    echo 0
