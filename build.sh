#!/bin/bash

export CGO_ENABLED=0
export GOOS=linux

PWD=(pwd)
cd "${0%/*}"

echo "Compiling..."
go build -a -installsuffix cgo -o main.exe .
echo "Building image..."
docker build -t oauth:latest .

cd $PWD
