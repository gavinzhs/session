#!/bin/bash
echo "remote restart"
ulimit -n 30000
fuser -k gotest
GOMAXPROCS=4 GOGC=50 nohup ./gotest &
