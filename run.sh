#!/usr/bin/env bash

# build
docker build -t kubia:v0.1 .

# run
docker run -d -p 8080:8080 --name kubia-container kubia:v0.1
