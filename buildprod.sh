#! /bin/bash

# Build web and other services

cd ~/go/src/YuanTube/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ~/go/src/YuanTube/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/go/src/YuanTube/stream
env GOOS=linux GOARCH=amd64 go build -o ../bin/stream

cd ~/go/src/YuanTube/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web