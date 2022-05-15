#!/bin/sh
echo "Checking for bin folder"
if [ ! -d "./bin" ]
then
    echo "bin folder not found... Creating..." 
    mkdir bin
fi
echo "Compiling new executable"
go build -o bin
echo "Running project"
./bin/jarvis-bot