#!/bin/sh
cd ~/JarvisBot/
echo "Checking for bin folder"
if [ ! -d "./bin" ]
then
    echo "bin folder not found... Creating..." 
    mkdir bin
fi
echo "Compiling new executable"
go build -o bin
echo "Starting nohup operation in background"
echo "Logs are stored in logfile.txt"
nohup ./bin/jarvis-bot > logfile.txt &