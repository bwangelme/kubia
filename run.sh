#!/usr/bin/env bash

# build
docker build -t bwangel/kubia:v0.4 .

# run
docker run -d -p 8080:8080 -p 8090:8090 --name kubia-container bwangel/kubia:v0.4
