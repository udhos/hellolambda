#!/bin/bash

GOOS=linux go build -o main
zip deployment.zip main
