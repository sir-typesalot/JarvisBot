#!/bin/sh
echo "Starting up app at $(date)"
name="app_screen"
echo "Compiling new executable"
cd ~/JarvisBot/
go build -o bin
screen -c log.conf -dm -S $name bin/jarvis-bot
echo "Restarted detached screen $name"