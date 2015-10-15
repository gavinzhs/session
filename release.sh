#!/bin/bash

host=192.168.1.200

dir=/opt/gotest
app=gotest

echo "cross compile"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

echo "prepare"
ssh root@$host "mkdir $dir"


echo "scp"
scp restart.sh root@$host:$dir/
scp $app root@$host:$dir/$app.new

echo "remote operate"
ssh root@$host "cd $dir && fuser -k $app"
ssh root@$host "cd $dir && mv $app.new $app && sh restart.sh `</dev/null` >nohup.out 2>&1 &"

rm -rf $app
echo "done"