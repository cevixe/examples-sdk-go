#!/usr/bin/env bash

export AWS_STACK=cevixe-3factor-example
export AWS_BUCKET=sam-poc-bucket
sam deploy --stack-name $AWS_STACK --s3-bucket $AWS_BUCKET --capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
