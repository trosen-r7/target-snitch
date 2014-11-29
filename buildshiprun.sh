#!/bin/bash

GOARCH=amd64 GOOS=linux go build
scp target-snitch trevor@192.168.56.2:~
ssh -t trevor@192.168.56.2 "./target-snitch"
