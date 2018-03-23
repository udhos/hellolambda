#!/bin/bash

gofmt -s -w *.go

GOOS=linux go build -o main

zip deployment.zip main
