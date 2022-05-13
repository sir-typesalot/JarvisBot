#!/bin/sh
echo "Last reboot time: $(date)" > /etc/motd
screen -dm -S "app_dev" ./jarvis-bot
echo "Restarted detached screen 'app_dev'"