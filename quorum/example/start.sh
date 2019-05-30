#!/bin/bash

ps aux | grep crux | awk '{print $2}'| xargs kill -9
ps aux | grep geth | awk '{print $2}'| xargs kill -9
./raft-init.sh
./raft-start.sh

tail -f /dev/null