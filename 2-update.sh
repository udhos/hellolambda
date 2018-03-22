#!/bin/sh

aws lambda update-function-code --function-name HelloFunction --zip-file fileb://deployment.zip
