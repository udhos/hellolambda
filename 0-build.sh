#!/bin/bash

go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/aws/session
go get github.com/aws/aws-sdk-go/service/s3

gofmt -s -w *.go

GOOS=linux go build -o main

zip deployment.zip main
