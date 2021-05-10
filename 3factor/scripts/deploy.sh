#!/usr/bin/env bash

while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -a|--application)
    CEVIXE_APP="$2"
    shift
    shift
    ;;
    -s|--stack)
    AWS_STACK="$2"
    shift
    shift
    ;;
    -b|--bucket)
    AWS_BUCKET="$2"
    shift
    shift
    ;;
esac
done

echo "CEVIXE APP = ${CEVIXE_APP}"
echo "AWS STACK  = ${AWS_STACK}"
echo "AWS BUCKET = ${AWS_BUCKET}"

GRAPHQL_SCHEMA="s3://$AWS_BUCKET/schemas/$(cat .aws-sam/schema/final.graphql | shasum | head -c 40).graphql" && \
aws s3 cp .aws-sam/schema/final.graphql "${GRAPHQL_SCHEMA}" && \
sam deploy --stack-name "${AWS_STACK}" --s3-bucket "${AWS_BUCKET}" --parameter-overrides ApplicationName="${CEVIXE_APP}" SchemaDefinition="${GRAPHQL_SCHEMA}" --capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
