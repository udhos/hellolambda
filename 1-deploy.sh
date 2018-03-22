#!/bin/bash

die () {
    echo 1>&2 $*
    exit 1
}

# create role
aws iam create-role --role-name lambda_basic_execution \
    --assume-role-policy-document '{
    "Statement": [{
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }]
  }'

# attach policy
aws iam attach-role-policy --role-name lambda_basic_execution --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

ROLE_ARN=`aws iam get-role --role-name lambda_basic_execution --query 'Role.Arn' --output text`

[ -n "$ROLE_ARN" ] || die "missing env var ROLE_ARN=arn:aws:iam::<account-id>:role/<role>"

echo ROLE_ARN=$ROLE_ARN

# create function
aws lambda create-function \
    --region us-east-1 \
    --function-name HelloFunction \
    --zip-file fileb://./deployment.zip \
    --runtime go1.x \
    --role "$ROLE_ARN" \
    --handler main

