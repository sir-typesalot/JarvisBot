#!/bin/sh
cd ~/JarvisBot/
echo "Compiling new executable"
go build -o bin
nohup ./bin/jarvis-bot > logfile.txt &