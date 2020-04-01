#!/usr/bin/env bash

# build
docker build -t bwangel/kubia:v0.2 .

# run
docker run -d -p 8080:8080 --name kubia-container bwangel/kubia:v0.2
