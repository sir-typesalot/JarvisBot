#!/bin/sh
# Kill current app process
echo "Killing running process"
kill -9 $(pgrep -f './bin/jarvis-bot')
cd ~/JarvisBot/
echo "Pulling new changes to project"
git stash && git pull
echo "Restarting bot"
sh startup.sh
