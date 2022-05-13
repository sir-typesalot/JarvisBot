#!/bin/sh
echo "Last reboot time: $(date)" > /etc/motd
name="app_dev"
path="/var/log";
config="logfile ${path}/${name}.log
logfile flush 1
log on
logtstamp after 1
logtstamp string \"[ %t: %Y-%m-%d %c:%s ]\012\"
logtstamp on";
echo "$config" > /tmp/log.conf
echo "Created log file"
screen -c /tmp/log.conf -dmSL "$name" ./jarvis-bot
rm /tmp/log.conf
echo "Restarted detached screen $name"